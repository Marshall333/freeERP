package main

import (
	"flag"
	"fmt"
	com "freeERP/free_erp/common"
	"freeERP/free_erp/handler"
	confUtil "freeERP/free_erp/utils/config"
	sqlLib "freeERP/free_erp/utils/mysql"
	"github.com/CodisLabs/codis/pkg/utils/log"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	soldManager handler.SoldManager
	repoManager handler.RepoManager
)

func InitServer(confPath string) error {
	com.GConfig = new(com.Config)
	if err := confUtil.LoadConfig(com.GConfig, confPath); err != nil {
		log.Errorf("InitServer confPath[%s] error[%v]", confPath, err)
		return err
	}
	com.GSQLHelper = new(sqlLib.SQLHelper)
	com.GSQLHelper.Init(com.GConfig.SQLAddr)

	return nil
}

func main() {
	confPath := flag.String("c", "", "")
	flag.Parse()
	if *confPath == "" {
		fmt.Println("config path can't null!")
		//log.Errorf("config path can't null!")
		return
	}

	InitServer(*confPath)

	com.Gin = gin.Default()

	//com.Gin.Use(GinLoggerWithConfig(os.Stdout))

	// 获取商品列表
	com.Gin.Any("/get_commodity_list", repoManager.GetCommodityList)

	// 获取订单列表
	com.Gin.Any("/get_order_list", soldManager.GetOrderList)
	// 获取订单详情
	com.Gin.Any("/get_order_details", soldManager.GetOrderDetails)
	// 下单
	com.Gin.Any("/create_order", soldManager.CreateOrder)

	com.Gin.Any("/login", handler.Login)
	com.Gin.Any("/logout", handler.Logout)
	com.Gin.Any("/get_auth", handler.GetAuth)

	com.Gin.Run(com.GConfig.ListenAddr)
}
