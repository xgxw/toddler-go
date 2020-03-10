# readme
rules 是 rule 工厂. 包含工厂类, 创造者模式, 以及工厂的各个方法.

`rules.go` 定义了Rule服务的接口.

`factory.go` 定义了工厂类的接口和Creator的接口.
- 工厂类提供了根据name创建具体类的方法(Get()). 其次, 在这里, 我实现了获取所有
  已实现的工厂类name集合 的方法(GetAllRuleNames()).
- Creator: 创造者模式. 每个Rule只要实现Creator, 就可以使用统一的方式创建实例.
  注意其中使用的 Resource 和如何实现的.

`demo_creator.go` 和 `demo.go` 是具体rule实现的示例, 工厂类使用参考
`internal/services/`
- 在这里, 我通过 init() 函数实现了将该工厂类添加到类集合中, 方便调用者获取
  所有已实现的工厂类.
