package main

import (
	"fmt"
	_ "github.com/zhuCheer/gateway/flag"
	"github.com/zhuCheer/gateway/http"
	_ "github.com/zhuCheer/gateway/model"
	"github.com/zhuCheer/gateway/proxy"
	"time"
)

func main() {
	fmt.Println("start smart proxy")

	go proxy.ProxyStart()

	// 延时2s启动api
	time.Sleep(2 * time.Second)
	http.ApiStart()

}
