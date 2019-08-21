package http

import "github.com/zhuCheer/gateway/http/action"

func registRoute() {
	httpSrv := GroupNode("/api", "api")

	httpSrv.GET("/reload", action.ReloadNodes)
	httpSrv.ALL("/ping", action.Ping)
}
