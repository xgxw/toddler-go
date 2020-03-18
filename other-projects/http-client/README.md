# HTTP Client
搭建Http客户端.

Http客户端主要分为两点: Client 和 业务方法.

Client 构建主要为
1. NewClient(): 创建客户端
2. NewRequest(): 创建一个Request. 包括构建URL, 添加路径参数, 构建body.
3. Do(): 发起请求, 构建Response, 处理超时与error.

参考 client.go, request.go, Response.go

---
业务上定义自己的结构体和业务方法即可, 然后调用 NewRequest和Do 实现基础功能.

具体参考 demo.go

