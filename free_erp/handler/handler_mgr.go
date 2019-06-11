package handler

import (
	//log "common/alog"
	"fmt"
	com "freeERP/free_erp/common"
	"freeERP/free_erp/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

const (
	KLoginTTL = 2 * 3600 // 登录有效期
)

var store = sessions.NewCookieStore([]byte("user_info"))

func init() {
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   24 * 3600,
		HttpOnly: true,
	}
}

// 获取user_id
func GetUserID(sessionid string) int64 {
	timestamp := time.Now().Unix()
	sqlClause := utils.GetEscapeSqlClause("select user_id from user where sessionid='%s' and session_ttl < %d", sessionid, timestamp)
	datas, err := com.GSQLHelper.GetQueryDataList(sqlClause)
	if err != nil || len(datas) == 0 {
		fmt.Printf("GetUserID sessionid[%s] error[%v] timestamp[%d] sqlClause[%s]", sessionid, err, timestamp, sqlClause)
		//log.Errorf("GetUserID sessionid[%s] error[%v] timestamp[%d] sqlClause[%s]", sessionid, err, timestamp, sqlClause)
		return -1
	}
	return utils.StrToInt64(datas[0]["user_id"])
}

// 登录
func Login(c *gin.Context) {
	username := utils.EscapeStringBackslash(c.Query("username"))
	password := utils.EscapeStringBackslash(c.Query("password"))
	fmt.Printf("Login username[%s] password[%s] IP[%s]", username, password, c.ClientIP())
	//log.Infof("Login username[%s] password[%s] IP[%s]", username, password, c.ClientIP())

	sqlClause := utils.GetEscapeSqlClause("select user_id, username, auth, sessionid, session_ttl from user where username='%s' and password='%s'", username, password)
	datas, err := com.GSQLHelper.GetQueryDataList(sqlClause)
	if err != nil || len(datas) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "login error!"})
		fmt.Printf("Login username[%s] password[%s] error[%v] sqlClause[%s]", username, password, err, sqlClause)
		//log.Errorf("Login username[%s] password[%s] error[%v] sqlClause[%s]", username, password, err, sqlClause)
		return
	}

	userInfo := make(map[string]interface{})
	userInfo["user_id"] = datas[0]["user_id"]
	userInfo["username"] = datas[0]["username"]
	userInfo["sessionid"] = datas[0]["sessionid"]
	if datas[0]["sessionid"] == "" {
		userInfo["sessionid"] = utils.GetMd5Value(fmt.Sprintf("%d%s%s", time.Now().UnixNano(), c.ClientIP(), userInfo["username"]))
	}

	com.GSQLHelper.UpdateDataByMap("user", map[string]interface{}{"sessionid": userInfo["sessionid"], "session_ttl": time.Now().Unix() + KLoginTTL},
		" where user_id="+datas[0]["user_id"])

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "user_info": userInfo})
}

func GetAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "auth": []string{"sold_out", "goods_info", "repository_info", "report"}})
}

func Logout(c *gin.Context) {
	username := c.Query("username")

	fmt.Printf("Logout username[%s] IP[%s]", username, c.ClientIP())
	//log.Infof("Logout username[%s] IP[%s]", username, c.ClientIP())

	com.GSQLHelper.UpdateDataByMap("user", map[string]interface{}{"sessionid": ""}, " where username='"+username+"'")

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": ""})
}
