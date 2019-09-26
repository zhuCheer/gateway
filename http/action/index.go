package action

import (
	"fmt"
	"github.com/zhuCheer/gateway/http/context"
	"github.com/zhuCheer/gateway/model"
	"github.com/zhuCheer/gateway/proxy"
	"log"
	"reflect"
)

func Ping(ctx *context.Context) {
	fmt.Println("ping2")
	ctx.Response().Write([]byte("pong"))
	return
}

func Test(ctx *context.Context) {
	fmt.Println("ping2")
	ctx.Json(map[string]interface{}{
		"code": 0,
		"msg":  "1",
		"data": "",
	})
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

	balanceType := ""
	switch reflect.TypeOf(info.Balancer).String() {
	case "*balancer.RoundRobinLoad":
		balanceType = "roundrobin"
	case "*balancer.RandomLoad":
		balanceType = "random"
	case "*balancer.WRoundRobinLoad":
		balanceType = "wraoudrobin"
	}

	if err != nil {
		ctx.Json(map[string]interface{}{
			"error": 1,
			"info":  err.Error(),
			"data":  "",
		})
		return
	}

	ctx.Json(map[string]interface{}{
		"error":       0,
		"info":        "success",
		"data":        info,
		"balanceType": balanceType,
	})
	return
}

// RegisterSite 注册一个站点
func RegisterSite(ctx *context.Context) {
	domain := ctx.ParamsPost("domain")
	balance := ctx.ParamsPost("balance")
	scheme := ctx.ParamsPost("scheme")
	if domain == "" {
		ctx.Json(map[string]interface{}{
			"error": 1,
			"info":  "domain/addr is empty",
			"data":  "",
		})
		return
	}
	fmt.Println("api regist domain addr", domain)
	proxy.ProxySrv.RegistSite(domain, balance, scheme)
	ctx.Json(map[string]interface{}{
		"error": 0,
		"info":  "success",
		"data":  "",
	})
	return
}

// InsertNode 新增一个节点
func InsertNode(ctx *context.Context) {
	domain := ctx.ParamsPost("domain")
	addr := ctx.ParamsPost("addr")
	weight := ctx.ParamsPostInt("weight")

	if domain == "" || addr == "" {
		ctx.Json(map[string]interface{}{
			"error": 1,
			"info":  "domain/addr is empty",
			"data":  "",
		})
		return
	}
	fmt.Println("api add domain addr", domain, addr)
	proxy.ProxySrv.DelAddr(domain, addr)
	proxy.ProxySrv.AddAddr(domain, addr, uint32(weight))
	ctx.Json(map[string]interface{}{
		"error": 0,
		"info":  "success",
		"data":  "",
	})
	return
}

// 删除一个节点
func RemoveNode(ctx *context.Context) {
	domain := ctx.ParamsPost("domain")
	addr := ctx.ParamsPost("addr")
	if domain == "" || addr == "" {
		ctx.Json(map[string]interface{}{
			"error": 1,
			"info":  "domain/addr is empty",
			"data":  "",
		})
		return
	}
	fmt.Println("api remove domain addr", domain, addr)
	proxy.ProxySrv.DelAddr(domain, addr)
	return
}

// 清空一个站点
func FlushSite(ctx *context.Context) {
	domain := ctx.ParamsPost("domain")
	if domain == "" {
		ctx.Json(map[string]interface{}{
			"error": 1,
			"info":  "domain is empty",
			"data":  "",
		})
		return
	}
	fmt.Println("api flush domain", domain)
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
	domain := ctx.ParamsPost("domain")
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
	if info.ID > 0 {
		proxy.ProxySrv.RegistSite(domain, info.Balance, info.Scheme)
		nodes := model.NodesDB.QueryListBySiteId(info.ID)
		for _, node := range nodes {
			proxy.ProxySrv.AddAddr(domain, node.Addr, node.Weight)
			log.Printf("└ proxy node %s %d", node.Addr, node.Weight)
		}
	}

	fmt.Println("api reload domain", domain)
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
