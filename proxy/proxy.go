package proxy

import (
	"github.com/zhuCheer/gateway/flag"
	"github.com/zhuCheer/gateway/model"
	"github.com/zhuCheer/libra"
	"log"
	"time"
)

var ProxySrv *libra.ProxySrv

// 开启代理服务
func ProxyStart() {
	// 延时2s启动
	time.Sleep(2 * time.Second)
	proxyBindAddr := flag.Config.GetString("app.proxy_addr")
	loggerLevel := flag.Config.GetString("app.logger_level")

	ProxySrv := libra.NewHttpProxySrv(proxyBindAddr, nil)
	list := model.SitesDB.GetAll()
	for _, item := range list {
		log.Printf("register proxy %s://%s %s", item.Scheme, item.Domain, item.Balance)
		ProxySrv.RegistSite(item.Domain, item.Balance, item.Scheme).SetLoggerLevel(loggerLevel)

		nodes := model.NodesDB.QueryListBySiteId(item.ID)
		for _, node := range nodes {
			ProxySrv.AddAddr(item.Domain, node.Addr, node.Weight)
			log.Printf("└ proxy node %s %d", node.Addr, node.Weight)
		}
	}

	// start proxy server
	ProxySrv.Start()
}
