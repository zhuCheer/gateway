package main

import (
	"fmt"
	_ "github.com/zhuCheer/gateway/flag"
	"github.com/zhuCheer/gateway/http"
	_ "github.com/zhuCheer/gateway/model"
	"github.com/zhuCheer/gateway/proxy"
)

func main() {
	fmt.Println("start smart proxy")

	go proxy.ProxyStart()
	http.ApiStart()

}
