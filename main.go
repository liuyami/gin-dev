package main

import (
	"hanya-gin/bootstrap"
	"hanya-gin/global"
)

func main() {

	// 初始化配置
	bootstrap.InitializeConfig()

	// 日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("Application initial...")

	// 数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 初始化验证器
	bootstrap.InitializeValidator()

	// 初始化 redis
	global.App.Redis = bootstrap.InitializeRedis()

	// 初始化 storage
	bootstrap.InitializeStorage()

	bootstrap.RunServer()
}
