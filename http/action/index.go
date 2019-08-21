package action

import (
	"fmt"
	"github.com/zhuCheer/gateway/http/context"
	"github.com/zhuCheer/gateway/model"
	"github.com/zhuCheer/gateway/proxy"
	"log"
)

func Ping(ctx *context.Context) {
	fmt.Println("ping2")
	ctx.Response().Write([]byte("pong"))
	return
}

// 刷新节点信息
func ReloadNodes(ctx *context.Context) {

	list := model.SitesDB.GetAll()
	for _, item := range list {
		log.Printf("fresh proxy %s://%s %s", item.Scheme, item.Domain, item.Balance)
		proxy.ProxySrv.FlushProxy(item.Domain)
		proxy.ProxySrv.RegistSite(item.Domain, item.Balance, item.Scheme)
		nodes := model.NodesDB.QueryListBySiteId(item.ID)
		for _, node := range nodes {
			proxy.ProxySrv.AddAddr(item.Domain, node.Addr, node.Weight)
			log.Printf("└ proxy node %s %d", node.Addr, node.Weight)
		}
	}

	ctx.Response().Write([]byte("done"))
}
