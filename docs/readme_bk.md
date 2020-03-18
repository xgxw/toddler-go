# toddler-go
学习使用go搭建不同类型的项目.

基于微服务, docs 文档帮助, 用于入门按阶段学习, 源码用于熟悉后直接查看. 如我自己使用本项目时.

## 开发
web/grpc/crontab 分别是
1. web: web服务. 入口: `internal/cmd/demo_server.go`
2. grpc: grpc服务. 入口: `internal/cmd/demo_grpc.go`
3. crontab: 离线脚本, 定时脚本. 入口: `internal/cmd/demo_task1.go`

other-projects 包括
1. httpclient. 一个 http服务的 sdk.

这里推荐一个使用go写的常用基础工具包[foundation-go](https://github.com/xgxw/foundation-go).
封装了常用的功能, 如 db,log 等.

---
项目介绍

项目使用 corba 构建cli程序, 并且支持多个入口, 比如web服务, grpc服务等.
参考 `internal/cmd/`

项目外层定义对外开放的数据, 如结构体, 服务接口. 参考 `demo.go`

项目外层同样定义了API接口. 对于 grpc 协议, 因为请求格式比较固定且多为内部调用,
所以 proto 文件基本可以描述接口格式和规范. 对于web服务, 借助 swagger 编写API文档.

internal 定义只能项目内部访问的数据. 其中, cmd 定义cli入口, codes 定义错误码,
controllers/middlewares/services 如其名字, 定义控制器/中间件/服务.
- go 仅允许项目内部代码访问 internal.

如果有只允许本项目访问的结构体, 应当定义到 `internal/models` 中.

`internal/tests` 定义了一些单元测试使用的工具

此外, 项目实现了一些其他常用功能
1. 常用方法, 即一些常用的方法示例, 如读写文件.
  - 参考 `utils`
2. 工厂模式, RuleService 通过调用 ruleFactory 的方法生产具体的类, 执行相应业务.
  - 参考 `internal/rules/`

## 部署
目前一般采用docker部署. 项目通过 Dockerfile/docker-entrypoint 构建Docker镜像.

Makefile 用于项目构建, 如编译代码, 也可以用于构建docker, 推送镜像, 编译proto等.
使用 makefile 可以减少我们需要重复记忆和操作的情况.

`script/deploy/run.sh` 是一个在服务端运行的脚本. 当使用非docker形式部署时(如在测试环境),
可以借助该脚本, 使用 supervisor 托管程序.

