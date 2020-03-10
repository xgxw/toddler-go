# readme
rules 是 rule 工厂.

实现了创造者模式, 工厂模式, 以及具体的工厂类(业务相关). 主要分为如下文件

`rules.go` 定义了Rule服务的接口, 提供统一的调用入口, 各工厂类实现自己具体的业务.

`factory.go` 定义了工厂模式的接口和Creator的接口.
- 工厂最核心的职责是Get()接口, 传入工厂类的Name, 返回该工厂类的实例. 其次,
  由于业务需要, 在这里我实现了allRuleCreators, 用于获取所有已实现的工厂类Names.
- Creator: 创造者模式. 每个Rule只要实现Creator, 就可以使用统一的方式创建实例.
  使用 Resource 加载所有工厂类需要使用的资源.

`demo.go` 负责具体的业务实现.

`demo_creator.go` 负责创建工厂类demo, 并且通过 init() 函数实现了将该工厂类添加到
  类集合中, 方便工厂获取所有已实现的工厂类.

工厂类使用参考 `internal/services/`