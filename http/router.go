package http

import "github.com/zhuCheer/gateway/http/action"

func registRoute() {
	httpSrv := GroupNode("/api", "api")

	httpSrv.GET("/reload", action.ReloadAllNodes)
	httpSrv.GET("/reloadone", action.ReloadSite)
	httpSrv.GET("/flushone", action.FlushSite)
	httpSrv.ALL("/ping", action.Ping)

	httpSrv.GET("/info", action.GetInfo)
}
