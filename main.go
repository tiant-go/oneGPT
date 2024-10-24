package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tiant-go/golib"
	"github.com/tiant-go/golib/flow"
	"oneGPT/conf"
	"oneGPT/helpers"
	"oneGPT/router"
)

func main() {
	// 1.全局变量初始化
	helpers.PreInit()
	defer helpers.Clear()
	// 2 启动器创建
	engine := gin.New()
	golib.Bootstraps(engine, conf.WebConf)
	// 3 初始化资源
	helpers.InitResource()
	// 4.初始化http服务路由
	router.Http(engine)
	// 5.框架启动
	flow.Start(engine, conf.WebConf, func(engine *gin.Engine) (err error) {
		flow.SetDefaultDBClient(helpers.MysqlClient)
		flow.SetDefaultRedisClient(helpers.RedisClient)
		return nil
	})
}
