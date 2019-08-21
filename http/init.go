package http

import (
	"context"
	"fmt"
	"github.com/zhuCheer/gateway/flag"
	httpctx "github.com/zhuCheer/gateway/http/context"
	"net/http"
)

type routerServer struct {
	prefix string
	name   string
}

type routerNode struct {
	group   string
	patten  string
	method  string
	handler func(ctx *httpctx.Context)
}

var routers []routerNode

// GET 注册 get 请求
func (r *routerServer) GET(patten string, handler func(ctx *httpctx.Context)) {
	routers = append(routers, routerNode{r.name, r.prefix + patten, "GET", handler})
}

// POST 注册 post 请求
func (r *routerServer) POST(patten string, handler func(ctx *httpctx.Context)) {
	routers = append(routers, routerNode{r.name, r.prefix + patten, "POST", handler})
}

// ALL 兼容所有请求
func (r *routerServer) ALL(patten string, handler func(ctx *httpctx.Context)) {
	routers = append(routers, routerNode{r.name, r.prefix + patten, "ALL", handler})
}

func DefaultNode() *routerServer {
	return &routerServer{}
}

func GroupNode(prefix, name string) *routerServer {
	return &routerServer{prefix, name}
}

func handlerFunc(node routerNode) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		method := request.Method
		if node.method != "ALL" && node.method != method {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte("Not Found"))
		} else {
			ctx := httpctx.New(context.Background(), writer, request)
			node.handler(ctx)
		}
	}
}

// 开启 api 接口服务
func ApiStart() {
	apiBindAddr := flag.Config.GetString("app.api_addr")
	httpApiMux := http.NewServeMux()

	// 注册路由
	registRoute()
	for _, item := range routers {
		if item.group != "api" {
			continue
		}

		httpApiMux.HandleFunc(item.patten, handlerFunc(item))
		fmt.Println("regist uri:" + item.patten)
	}

	server := &http.Server{
		Addr:    apiBindAddr,
		Handler: httpApiMux,
	}
	fmt.Println("start listen " + apiBindAddr + " for api http server.")
	server.ListenAndServe()
}
