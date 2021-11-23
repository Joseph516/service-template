# Service-Template

## Quick start

1. 进入service-template目录下，建立数据库与表：

```shell
# 在configs/config.yaml中配置自己的设备信息
cd service-template && bash scripts/database.sh
```

2. 在service-template目录下，启动后端：

```shell
go run main.go
```

## 目录结构

```shell
$ user in /path/to/service-template
├── README.md
├── bin
├── configs # 配置文件
│   ├── config.yaml
│   └── service.sql
├── global # 全局变量
│   ├── db.go
│   └── setting.go
├── go.mod
├── go.sum
├── internal # 内部模块
│   ├── dao # 数据访问层，具体操作数据库的地方
│   ├── middleware # Http中间层
│   ├── model # 模型层，用户存放model文件
│   ├── routers # 路由相关的逻辑
│   └── service # 项目核心业务逻辑，每次操作时新建service对象。即：router->service->dao->model
├── main.go
├── pkg # 项目相关的模块包
│   ├── app
│   ├── database
│   ├── errcode
│   ├── settings
│   ├── upload
│   └── util
├── scripts # 各类构建、安装、分析等操作脚本
│   └── database.sh
└── storage # 文件上传存储位置
```

## TODO

1. Swagger生成web api。