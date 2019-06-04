package handler

import (
	com "freeERP/free_erp/common"
	"net/http"
	"time"

	"github.com/CodisLabs/codis/pkg/utils/log"
	"github.com/gin-gonic/gin"
)

type RepoManager struct {
}

// 获取商品列表
func (rm *RepoManager) GetCommodityList(c *gin.Context) {
	sqlClause := "select goods_id,commodity_id,name,`describe`,out_price,remark from commodity_info;"

	datas, err := com.GSQLHelper.GetQueryDataList(sqlClause)
	if err != nil {
		log.Errorf("RepoManager GetCommodityList error[%v] sqlClause[%s]", err, sqlClause)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "datas": datas})
}

// 新建商品
func (rm *RepoManager) CreateNewCommodity(c *gin.Context) {
	data := make(map[string]interface{})

	data["commodity_id"] = c.Query("commodity_id")
	data["name"] = c.Query("name")
	data["describe"] = c.Query("describe")
	data["in_price"] = c.GetFloat64("in_price")
	data["out_price"] = c.GetFloat64("out_price")

	data["createtime"] = time.Now().Format(com.KTimeFormat)
	data["status"] = 1

	rows, err := com.GSQLHelper.InsertDataByMap("commodity_info", data)
	if err != nil || rows <= 0 {
		log.Errorf("RepoManager CreateNewCommodity data[%v] err[%v] rows[%d]", data, err, rows)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "创建商品失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": ""})
}
