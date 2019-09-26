# һ��api���ط���ʵ��


## �������
- [libar](github.com/zhuCheer/libra) ����һ������ʵ��
- ��������ͨ�� Mysql ���й���ά��
- ֧�ֹ����� api �ӿ�
- ͨ�� UI ���ʵ�ֶ�Ӧ��ʵ�ֿ��ٵľ��������ش


## ��ʼʹ��

1. `git clone https://github.com/zhuCheer/gateway`
2. �������������,������ golang; �汾��1.11����ǰ����Ҫ��ִ����`export GO111MODULE=on` ���� gomudule
Ȼ��ͨ��������������������;
 ```
go mod tidy
go mod vendor
```
3.�� gateway.sql ���뵽 MySql;
4.�༭ `config/config.toml` ����mysql����
5. �������� `go run main.go --config=config/config.toml`


## ����˵��

- ��������������������˿� Ĭ�� 5000 �˿ں� 5001 �˿�;
- �˿ںſ�����`config/config.toml`�н�������
- `proxy_addr` �����ش�������˿�Ĭ��5000;
- `api_addr` �ǹ�����api�ӿڷ���˿�,Ĭ��5001;

## ���ṹ˵��
- `qi_sites` վ���,����վ�����������ؾ������͵���Ϣ;
- `qi_nodes` �ڵ��������վ���¶�Ӧ�Ļ���ip:port�Լ�Ȩ����Ϣ,�� `site_id` �ֶ��� `qi_sites` �е� id ��Ӧ;


## �ܹ�ʵ��˵��

����ʵ�ֻ����Ͽ���ͨ����ͼ����ͣ�Ŀǰ������Ϊһ�������ʵ��, �� mysql �е�����-�ڵ����ݽ��д���, �û������������������� ip ��, ������������ͳһ��ת��;

![image](https://note.youdao.com/favicon.ico)

## �����˽ӿ�

- ���Բ�ѯվ���½ڵ���Ϣ���������͵�;

|��ѯָ��վ����Ϣ�ӿ�||
|----------|-----------|
| url      | http://127.0.0.1:5001/api/info|
| �������� | GET       | 


|����||
|----------|-----------|
| domain      | www.qiproxy.cn|

--------------

- ˢ��վ���������,�ڵ���Ϣ;

|ˢ��վ��ӿ�||
|----------|-----------|
| url      | http://127.0.0.1:5001/api/reloadone|
| �������� | POST       | 


|����||
|----------|-----------|
| domain      | www.qiproxy.cn|


