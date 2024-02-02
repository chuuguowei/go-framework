# go-framework
基于gin框架搭建的一套go应用web架构

## 目录结构

- bootstrap
- config
- controller
- library
- logs
- model
- .env.yaml
- go.mod
- main.go

## bootstrap
服务依赖初始化，包含日志、MySQL、Redis、Cache等，可自定义扩展

## config
集成gopkg.in/yaml.v3，用于解析yaml文件的配置项

## controller
业务入口层，接口API在这里定义和实现

## library
框架层
- cache - 本地缓存，集成go-cache
- client - 全局服务client
- common - 公共结构体和工具类封装
- config - yaml.v3配置实现
- logger - 日志打印实现，集成uber
- middleware - 中间件
- mysql - mysql存储实现底层封装
- redis - redis缓存封装
- route - 路由配置

## logs
日志文件

## model
业务实现层
- dao 数据库dao层封装（增删改查）
- service 业务实现封装
- types 业务模型封装

# 快速启动
* 将.env.yaml配置中DB和Redis配置替换为自己的服务配置
* 执行 go mod tidy 初始化项目依赖
* 执行 go run main.go 启动服务




