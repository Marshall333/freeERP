package common

import (
	sqlLib "freeERP/free_erp/utils/mysql"
	"github.com/gin-gonic/gin"
)

const (
	KTimeFormat = "2006-01-02 15:04:05"
	KDateFormat = "2006-01-02"
)

type Config struct {
	SQLAddr    string `json:"sql_addr"`
	ListenAddr string `json:"listen_addr"`
}

var (
	Gin *gin.Engine

	GSQLHelper *sqlLib.SQLHelper

	GConfig *Config
)
