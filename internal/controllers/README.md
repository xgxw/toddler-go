# controllers
控制层, 一般用于 web/grpc 路由参数的判断和解析, 封装 service 处理结果然后返回, 不负责具体的业务逻辑.

Controller 不负责业务逻辑, 业务逻辑由Service负责. 原因如下
1. 业务逻辑可能被多个服务, 多个Controller调用, 所以业务逻辑独立拆分到Service中.
2. Controller 作为控制器, 连通 web/grpc 服务与业务逻辑处理, 本身职责只有 参数校验和解析, 结果封装和返回.

参考 DemoController 中的 DoSomething 和 DoSomrthingHttp 方法, 就是对 Service.DoSomething 的封装.
Controller 之负责参数/结果的处理, 具体业务逻辑由 Service 负责.

Controller 也符合单一职责原则/领域模型, 一个 controller 负责一个service, 按业务划分Controller范围.
如 DemoController 和 DemoService, UserController 和 UserService.

---
当服务用到多个功能模块时, 如 UserService 同时需要访问 user 和 file(获取用户的头像信息) 时, 将 FileSvc
注入到 UserSvc 中, 而不是在 Controller 中定义多个 Service.

有些项目会添加 busnisee 层, 用于处理各 Service 的依赖. 即将各Service返回的结果在 business 层再处理一遍,
然后返回.
