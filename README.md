# 一个api网关服务实现


## 相关特性
- [libar](github.com/zhuCheer/libra) 包的一个典型实现
- 代理服务通过 Mysql 进行管理维护
- 支持管理端 api 接口
- 通过 UI 面板实现对应用实现快速的均衡器网关搭建

## 开始使用

1. `git clone https://github.com/zhuCheer/gateway`
2. 下载相关依赖包,如果你的 golang; 版本是1.11及以前的需要先执行下`export GO111MODULE=on` 开启 gomudule
然后通过下面命令下载依赖包;
 ```
go mod tidy
go mod vendor
```
3.将 gateway.sql 导入到 MySql

4.编辑 `config/config.toml` 配置mysql连接

5. 启动服务 `go run main.go --config=config/config.toml`

6. 搭建客户端UI可视化管理服务移步 [gatewayui](https://github.com/zhuCheer/gatewayui)



## 配置说明

- 服务启动后会启动两个端口 默认 5000 端口和 5001 端口;
- 端口号可以在`config/config.toml`中进行配置
- `proxy_addr` 是网关代理服务端口默认5000;
- `api_addr` 是管理端api接口服务端口,默认5001;

## 表结构说明
- `qi_sites` 站点表,保存站点域名，负载均衡类型等信息;
- `qi_nodes` 节点表，保存站点下对应的机器ip:port以及权重信息,有 `site_id` 字段与 `qi_sites` 中的 id 对应;


## 架构实现说明

网关实现基本上可以通过下图类解释，目前网关作为一个反向代实现, 将 mysql 中的域名-节点数据进行代理, 用户将域名都解析到网关 ip 上, 由网关来进行统一的转发;

![image](https://raw.githubusercontent.com/zhuCheer/gateway/master/flow.jpg)

## 管理端接口

#### 查询指定站点信息接口

> 可以查询站点下节点信息，均衡类型等;

|||
|----------|-----------|
| url      | http://127.0.0.1:5001/api/info|
| 请求类型 | GET       | 


|参数||
|----------|-----------|
| domain      | www.qiproxy.cn|

--------------

#### 刷新站点接口
> 刷新站点均衡类型,节点信息;

|||
|----------|-----------|
| url      | http://127.0.0.1:5001/api/reloadone|
| 请求类型 | POST       | 


|参数||
|----------|-----------|
| domain      | www.qiproxy.cn|



--------------

#### 新增节点接口
|||
|----------|-----------|
| url      | http://127.0.0.1:5001/api/insertone|
| 请求类型 | POST       | 


|参数|示例|说明|
|----------|-----------|-----------|
| domain      | www.qiproxy.cn| 域名|
| addr      | 192.168.1.101:80|节点ip端口|
| weight      | 10|权重|


--------------


#### 删除节点接口
|||
|----------|-----------|
| url      | http://127.0.0.1:5001/api/removenoe|
| 请求类型 | POST       | 


|参数|示例|说明|
|----------|-----------|-----------|
| domain      | www.qiproxy.cn| 域名|
| addr      | 192.168.1.101:80|节点ip端口|

