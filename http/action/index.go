package action

import (
	"fmt"
	"github.com/zhuCheer/gateway/http/context"
)

func Ping(ctx *context.Context) {
	fmt.Println("ping2")
	ctx.Response().Write([]byte("pong"))
	return
}

// 刷新节点信息
func ReloadNodes(ctx *context.Context) {
	fmt.Println("done1")
	ctx.Response().Write([]byte("done"))
}
