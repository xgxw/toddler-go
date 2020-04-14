# bootstrap
bootstrap 用于提供统一的入口, 来实例化系统的各种资源, 从而将资源实例化逻辑与其他逻辑解耦.

资源是指 db/cache 等基础资源, 以及 service/controller 等项目资源.

借助了如下手段实现
1. 单一职责原则. 资源实例化逻辑一其他逻辑解耦, 各自负责自己的部分.
2. 门面模式. 对外屏蔽实现. 当资源创建方法变化后, 只需要更改Bootstrap中的创建方法, 不需要改动其他地方.
3. 依赖注入/控制反转. 使用控制反转解决服务间的依赖, 而非内部解决或者使用单例.

基于Bootstrap类统一提供入口有如下优点
1. 低耦合, 高内聚.
2. 使代码更简洁. 提取出统一的创建方法可以减少代码重复度, 让代码更简洁直观.

## 实现
实现模板 [bootstrap example](/internal/cmd/bootstrap.go)

下面讲些题外话, 即bootstrap格式的发展历程

刚开始时, 是以成员字段的方式提供资源, 代码示例如下
```Go
type Bootstrap struct{
  Logger *flog.Logger
  DB     *database.DB
}
boot, teardown := newBootstrap(opts)
```

Command 通过调用 `boot.Logger` 方式获取资源.

但是这么做在某些情况下会有问题, 如
1. server/grpc 需要 logger/db/cache/demoSvc
2. task1 只需要 db
3. task2 只需要 demoSvc

使用这种方式初始化时, 在task1执行时不可避免的要创建 cache/demoSvc 实例, 但是程序不会用到.

所以现在改为 `GetLogger()` 的方式提供资源. 这种方式更灵活, 更能适应多个Command的情况.

使用这种方式可以更灵活的初始化, 有效避免task1运行时自动创建实例 cache/demoSvc.

## 误区
### 借助控制反转
所有的 controller/service 等初始化均应由外部控制, 避免内部初始化, 避免使用单例模式.

避免使用单例是因为单例模式在一定程度上违反了单一职责原则. 单例模式会导致 NewXX 函数包含其他的逻辑.
其次, 单例可以通过 依赖注入(或者说控制反转) 的方式实现, 所以应当避免使用单例模式.

如 UserController 依赖 UserService, 则应该在初始化 UserController 时初始化 UserService,
而非在 UserController 内部初始化 UserService. 伪代码如下
```Go
// 正确
func NewUserController(userSvc toddler.UserService)*UserController{
  return &UserController{
    userSvc: userSvc,
  }
}
// 错误
func NewUserController()*UserController{
  return &UserController{
    userSvc: services.NewUserService(),
  }
}
// 错误方式二: app提供统一的创建, app 作用类似 bootstrap
func NewUserController()*UserController{
  return &UserController{
    userSvc: app.NewUserService(),
  }
}
```

个人认为内部初始化方式不好是因为其 不符合单一职责, 增加了耦合度.
1. 增加了耦合度. 考虑如下使用情况. 当有多个UserService时(如v1/v2, 或不同业务逻辑的UserService,
  如 GoodUserService, OrderUserService), 调用 UserController 的调用者无法控制使用那个UserService.
2. NewUserController 的职责不单一, 其不仅负责创建 UserController, 还负责创建 UserService,
  使代码阅读性降低. controller/service 本来就应该是独立的一部分, 不应该掺杂太多其他的逻辑.

通俗来讲, 就是解耦解耦再解耦, 高内聚低耦合, 越是如此, 之后的维护和开发也就更容易.

### 内部定义
之前还见过一些其他组织结构的实现方式
````
- app/
- config.go
````

在 config.go 中定义 配置结构体, 在 app 中实例化资源. 然后在 Command 初始化时, 调用 `app.XX` 获取资源.
app.go 伪代码如下

```Go
type AppEntity struct{
  Logger *flog.Logger
  DB     *database.DB
}
var App *AppEntity
func init(){
  App = &AppEntity{}
}
func (app *AppEntity)GetLogger()*flog.Logger{...}
```

我个人认为这么写是不好的, 原因是 `如无必要, 勿增实体`
1. 领域划分不和是. 配置和实例创建过程都是项目私有的, 是在Command执行时初始化的, 应当是 Command 的一部分,
  不应该也没必要暴露给外部.
2. 增加了 `init()` 函数, 也没必要, 并且增加了不确定性(这个可以通过改为 `NewAppEntity()` 改善)
