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

func GetInfo(ctx *context.Context) {
	domain := ctx.ParamsGet("domain")

	if domain == "" {
		ctx.Json(map[string]interface{}{
			"error": 1,
			"info":  "domain is empty",
			"data":  "",
		})
		return
	}
	info, err := proxy.ProxySrv.GetSiteInfo(domain)

	if err != nil {
		ctx.Json(map[string]interface{}{
			"error": 1,
			"info":  err.Error(),
			"data":  "",
		})
		return
	}

	ctx.Json(map[string]interface{}{
		"error": 0,
		"info":  "success",
		"data":  info,
	})
	return
}

// 清空一个站点
func FlushSite(ctx *context.Context) {
	domain := ctx.ParamsGet("domain")
	if domain == "" {
		ctx.Json(map[string]interface{}{
			"error": 1,
			"info":  "domain is empty",
			"data":  "",
		})
		return
	}

	proxy.ProxySrv.FlushProxy(domain)
	ctx.Json(map[string]interface{}{
		"error": 0,
		"info":  "success",
		"data":  "",
	})
	return
}

// 刷新单个站点
func ReloadSite(ctx *context.Context) {
	domain := ctx.ParamsGet("domain")
	if domain == "" {
		ctx.Json(map[string]interface{}{
			"error": 1,
			"info":  "domain is empty",
			"data":  "",
		})
		return
	}

	info := model.SitesDB.QueryOneByDomain(domain)
	proxy.ProxySrv.FlushProxy(domain)
	proxy.ProxySrv.RegistSite(domain, info.Balance, info.Scheme)
	nodes := model.NodesDB.QueryListBySiteId(info.ID)
	for _, node := range nodes {
		proxy.ProxySrv.AddAddr(domain, node.Addr, node.Weight)
		log.Printf("└ proxy node %s %d", node.Addr, node.Weight)
	}

	ctx.Json(map[string]interface{}{
		"error": 0,
		"info":  "success",
		"data":  "",
	})
	return
}

// 刷新所有节点信息
func ReloadAllNodes(ctx *context.Context) {

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
