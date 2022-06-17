package config

import (
	"xorm.io/xorm"
	"xorm.io/xorm/names"
	// mysql 加载mysql
	_ "github.com/go-sql-driver/mysql"
)

// RequestIDKey ..
var RequestIDKey = "Th-Request-Id"

// DataAlreadyExist ..
var DataAlreadyExist uint16 = 1062

// ConfDB 是
var ConfDB *xorm.Engine

// InitDB 初始化 DB 连接，创建 xorm 的 engine，需要在服务启动访问 DB 前先调用
func InitDB() error {
	columnMapper := names.NewPrefixMapper(names.GonicMapper{}, "f_")
	tableMapper := names.NewPrefixMapper(names.GonicMapper{}, "t_")
	initEngine := func(engine *xorm.Engine) {
		engine.SetTableMapper(tableMapper)
		engine.SetColumnMapper(columnMapper)
	}
	engine, err := xorm.NewEngine("mysql", GetConf().DB.Mysql.Config.DSN)
	initEngine(engine)
	ConfDB = engine
	return err
}
