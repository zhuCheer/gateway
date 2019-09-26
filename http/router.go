package http

import "github.com/zhuCheer/gateway/http/action"

func registRoute() {
	httpSrv := GroupNode("/api", "api")

	httpSrv.POST("/reload", action.ReloadAllNodes)
	httpSrv.POST("/reloadone", action.ReloadSite)
	httpSrv.POST("/flushone", action.FlushSite)

	// 注册站点
	httpSrv.POST("/regist", action.RegisterSite)

	// 增加节点
	httpSrv.POST("/insertone", action.InsertNode)

	// 删除节点
	httpSrv.POST("/removenoe", action.RemoveNode)

	httpSrv.ALL("/ping", action.Ping)
	httpSrv.GET("/info", action.GetInfo)

	httpSrv.POST("/test", action.Test)
}
