# Toddler-go

go 项目学习, 示例.

为了写出更优雅, 结构更好的代码!

文件介绍如下
1. demo.proto: grpc API doc.
2. swagger.json: web api doc.
3. internal/cmd/demo_grpc.go: grpc 服务示例
4. internal/cmd/demo_task1.go: cli脚本示例
5. internal/rules 是 rule 工厂. 包含工厂类, 创造者模式, 以及工厂的各个方法.
6. 外层的 Makefile/Dockerfile/docker-entrypoint 是用于运维部署的.

demoService 是通过调用 ruleFactory 的方法执行业务的.
