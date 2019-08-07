# tom_club

## 环境变量
```
// 如果需要翻墙，还需要代理
export http_proxy=
export https_proxy=
```

## 
```
json:"-"    // 该 tag 序列化时忽略该字段
```

## 签名兼容 urlbase64 和 hex

## 最新步骤 1551852847
输出可执行文件
`GO111MODULE=on go build -o club_api_env/project/tom_club`

`cp config.ini.example club_api_env/project/config.ini`

`vim club_api_env/project/config.ini`

`vim club_api_env/.env`

`cd club_api_env`

`docker-compose up -d --build`

## 新加定时任务 1551951752

重启时也会先设置以下日志文件输出
新加定时任务，每天零点执行任务，重新生成日志文件。

输出可执行文件

```shell
GO111MODULE=on go build -o

cp club-prd club_api_env/project/worldmap-api

cp config.ini club_api_env/project/config.ini

cp -r updates/ club_api_env/project/updates

cp club_api_env/.env.example club_api_env/.env

cp LocList.xml.zh-cn ./club_api_env/project/

cp LocList.xml.en ./club_api_env/project/

vim club_api_env/project/config.ini 此处填写用于运行可执行文件的配置

vim club_api_env/.env 此处填写运行 docker 的配置

cd club_api_env

docker-compose up -d --build
```

交叉编译
`GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o club-prd`

https://www.jianshu.com/p/4b345a9e768e

$ docker logs -f -t CONTAINER_ID

https://www.jianshu.com/p/1eb1d1d3f25e

查看端口是否占用
netstat -an | grep 10205

## 数据库
### 初始化数据库

``  CREATE DATABASE IF NOT EXISTS `tom_club` DEFAULT CHARSET utf8 COLLATE utf8_general_ci; ``

### 删除数据库

`` DROP DATABASE `tom_club` ``

### 迁移数据库

`"github.com/golang-migrate/migrate"`

默认迁移所有文件 --help查看具体请求  迁移会在数据库建立迁移记录表，mysql参数记得带引号，不然会报错

` migrate -source file:// -database "mysql://root:ls950322@tcp(127.0.0.1:3306)/tom_club" up `

### 
经过测试，可以采用自动迁移的方式 -- 添加字段自动添加、修改和删除字段手动删除

按理说 属性一样的两个结构体应该可以视为同一个结构体。但是此处没成功。类型转换的方法暂未找到。
此处暂定model为最高级的包，即不从model 中导入其他的包，只在其他的包中引用model。