package flag

import (
	"github.com/zhuCheer/cfg"
	"log"
	"flag"
)

var Config = cfg.New("./config/config.toml")

func init(){
	log.Printf("init flag")

	path := flag.String("config", "./config/config.toml", "config file path")
	flag.Parse()
	Config = cfg.New(*path)

}
