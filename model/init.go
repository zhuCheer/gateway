package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhuCheer/gateway/flag"
	"time"
)


var (
	DB *sql.DB
	SitesDB *Sites
	NodesDB *Nodes
)
func init(){
	fmt.Println("init model")
	mysqlInit()

	SitesDB = new(Sites)
	NodesDB = new(Nodes)
}

func mysqlInit(){
	addr := flag.Config.GetString("database.addr")
	port := flag.Config.GetInt("database.port")
	username := flag.Config.GetString("database.username")
	password := flag.Config.GetString("database.password")
	dbname := flag.Config.GetString("database.dbname")
	maxConn := flag.Config.GetInt("database.connection_max")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",username,password,addr,port,dbname)
	conn,err := sql.Open("mysql",dsn)

	if err != nil{
		panic("Open mysql failed,err:"+err.Error())
		return
	}

	conn.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
	conn.SetMaxOpenConns(maxConn)//设置最大连接数
	conn.SetMaxIdleConns(16) //设置闲置连接数
	//验证连接
	if err := conn.Ping(); err != nil{
		fmt.Println("opon database fail")
		return
	}

	DB = conn
}