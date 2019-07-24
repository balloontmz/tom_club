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
`GO111MODULE=on go build -o email_api_env/project/tom_club`
`cp config.ini.example email_api_env/project/config.ini`
`vim email_api_env/project/config.ini`
`vim email_api_env/.env`
`cd email_api_env`
`docker-compose up -d --build`

## 新加定时任务 1551951752

重启时也会先设置以下日志文件输出
新加定时任务，每天零点执行任务，重新生成日志文件。

输出可执行文件

```shell
GO111MODULE=on go build -o

cp worldmap-prd worldmap_api_env/project/worldmap-api

cp config.ini worldmap_api_env/project/config.ini

cp -r updates/ worldmap_api_env/project/updates

cp worldmap_api_env/.env.example worldmap_api_env/.env

cp LocList.xml.zh-cn ./worldmap_api_env/project/

cp LocList.xml.en ./worldmap_api_env/project/

vim worldmap_api_env/project/config.ini 此处填写用于运行可执行文件的配置

vim worldmap_api_env/.env 此处填写运行 docker 的配置

cd worldmap_api_env

docker-compose up -d --build
```

交叉编译
`GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o worldmap-prd`

https://www.jianshu.com/p/4b345a9e768e

$ docker logs -f -t CONTAINER_ID

https://www.jianshu.com/p/1eb1d1d3f25e

查看端口是否占用
netstat -an | grep 10205