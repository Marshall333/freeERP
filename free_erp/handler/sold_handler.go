package handler

import (
	"fmt"
	com "free_erp/common"
	"free_erp/utils"
	"net/http"
	"strings"

	"github.com/CodisLabs/codis/pkg/utils/log"

	"github.com/gin-gonic/gin"
)

type SoldManager struct {
}

// 获取订单列表
func (sm *SoldManager) GetOrderList(c *gin.Context) {
	beginTime, _ := c.GetQuery("begin_time")
	endTime, _ := c.GetQuery("end_time")
	pageRow, _ := c.GetQuery("page_row")
	page, _ := c.GetQuery("page")

	pageRowNum := utils.StrToInt(pageRow)
	pageNum := utils.StrToInt(page)

	sqlClause := fmt.Sprintf("select * from `order_list` where createtime >= '%s' and createtime <= '%s' and status=1 limit %d,%d",
		beginTime, endTime, pageRowNum*(pageNum-1), pageRowNum)

	datas, err := com.GSQLHelper.GetQueryDataList(sqlClause)
	if err != nil {
		log.Errorf("SoldManager GetOrderList error[%v] sqlClause[%s]", err, sqlClause)
		c.JSON(http.StatusOK, map[string]interface{}{"code": -1, "msg": "查询失败"})
		return
	}
	resp := make(map[string]interface{})
	resp["code"] = 0
	resp["msg"] = ""
	resp["datas"] = datas
	c.JSON(http.StatusOK, resp)
}

// 获取订单详情
func (sm *SoldManager) GetOrderDetails(c *gin.Context) {
	orderID := c.Query("order_id")

	sqlClause := utils.GetEscapeSqlClause("select order_details_id, order_id, a.createtime, goods_num, commodity_id,name, total "+
		"from order_details as a left join commodity_info as b on a.goods_id=b.goods_id where order_id=%s and status=1", orderID)

	datas, err := com.GSQLHelper.GetQueryDataList(sqlClause)
	if err != nil {
		log.Errorf("SoldManager GetOrderDetails error[%v] sqlClause[%s]", err, sqlClause)
		c.JSON(http.StatusOK, map[string]interface{}{"code": -1, "msg": "查询失败"})
		return
	}
	resp := make(map[string]interface{})
	resp["code"], resp["msg"] = 0, ""
	resp["datas"] = datas
	c.JSON(http.StatusOK, resp)
}

// 根据商品ID查询商品信息
func GetCommodityInfoByID(comIDList []string) map[string]map[string]string {
	ret := make(map[string]map[string]string, 0)

	sqlClause := utils.GetEscapeSqlClause("select goods_id, commodity_id, name, describe, createtime,in_price,out_price where status=1 and goods_id in (%s)",
		strings.Join(comIDList, ","))
	datas, err := com.GSQLHelper.GetQueryDataList(sqlClause)
	if err != nil || len(datas) == 0 {
		log.Errorf("GetCommodityInfoByID error[%v] comIDList[%v] dataslen[%d] sqlClause[%s]", err, comIDList, len(datas), sqlClause)
		return ret
	}
	for _, data := range datas {
		ret[data["goods_id"]] = data
	}
	return ret
}

// 创建订单
func NewOrder(userID int64, goodsNumTab map[string]int, goodsInfo map[string]map[string]string) int64 {
	if len(goodsNumTab) == 0 || len(goodsInfo) == 0 || len(goodsInfo) < len(goodsNumTab) {
		return -1
	}

	totalMoney := 0.0
	for goodsID, num := range goodsNumTab {
		totalMoney += float64(num) * utils.StrToFloat64(goodsInfo[goodsID]["out_price"])
	}
	tx, err := com.GSQLHelper.GetSQLDB().Begin()
	if err != nil {
		log.Errorf("NewOrder get context error[%v]", err)
		return -1
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			log.Errorf("NewOrder recover panic[%v]", p)
		}
	}()
	result, err := tx.Exec(utils.GetEscapeSqlClause("insert into order_list(createtime,create_user_id,total,status) values(now(),%d,%f,2)", userID, totalMoney))
	if err != nil {
		log.Errorf("NewOrder insert order error[%v]", err)
		tx.Rollback()
		return -1
	}
	orderID, err := result.LastInsertId()
	if err != nil {
		log.Errorf("NewOrder get order id error[%v]", err)
		tx.Rollback()
		return -1
	}
	sqlClause := "insert into order_details(order_id,createtime,goods_id,goods_num,total) values"
	for goodsID, num := range goodsNumTab {
		sqlClause += utils.GetEscapeSqlClause("(%d,now(),%s,%d,%f),", orderID, goodsID, num, float64(num)*utils.StrToFloat64(goodsInfo[goodsID]["out_price"]))
	}
	sqlClause = strings.TrimRight(sqlClause, ",")
	if _, err = tx.Exec(sqlClause); err != nil {
		log.Errorf("NewOrder error[%v] sqlClause[%s]", err, sqlClause)
		tx.Rollback()
		return -1
	}
	tx.Commit()

	return orderID
}

// 下单
func (sm *SoldManager) CreateOrder(c *gin.Context) {
	sessionid := c.Query("sessionid")
	commodityList := c.Query("commodity_list") // 格式: 数量*商品ID,数量*商品ID.....

	userID := GetUserID(sessionid)
	if userID < 0 {
		log.Errorf("SoldManager CreateOrder sessionid[%s] login out of ttl", sessionid)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "请先登录"})
		return
	}

	comIDList := make([]string, 0)
	goodsNumTab := make(map[string]int, 0)

	isok := true

	commodityList = strings.Replace(commodityList, " ", "", -1)
	itemList := strings.Split(commodityList, ",")
	for _, item := range itemList {
		tmp := strings.Split(item, "*")
		num := tmp[0]
		goodsID := tmp[1]
		if !(len(tmp) == 2 && utils.StrToInt(num) > 0 && utils.StrToInt64(goodsID) > 0) {
			isok = false
			break
		}
		comIDList = append(comIDList, goodsID)
		goodsNumTab[goodsID] = utils.StrToInt(num)
	}
	if !isok {
		log.Errorf("SoldManager CreateOrder sessionid[%s] commodityList format error[%s]", sessionid, commodityList)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数格式错误"})
		return
	}

	goodsInfo := GetCommodityInfoByID(comIDList)
	if len(goodsInfo) != len(comIDList) {
		log.Errorf("SoldManager CreateOrder sessionid[%s] commodity id error[%s]", sessionid, commodityList)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "商品ID错误"})
		return
	}
	orderID := NewOrder(userID, goodsNumTab, goodsInfo)
	if orderID < 0 {
		log.Errorf("SoldManager CreateOrder sessionid[%s] com list[%s] New order", sessionid, commodityList)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "订单创建失败"})
		return
	}
	log.Infof("SoldManager CreateOrder sessionid[%s] commodityList[%s] orderID[%d]", sessionid, commodityList, orderID)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "order_id": orderID})
}
