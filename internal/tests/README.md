# tests

一些测试所用的mock

`db.go` 实现了数据库连接的mock和 `foundation-go/database.DB` 的模拟.
`mocks/` 文件夹是使用gomock生成的外层接口的mock文件. 生成方法参考 Makefile

具体的使用示例参考 [使用示例](/internal/services/demo_test.go)

常用命令
```Bash
# 执行特定的单元测试.
go test -v -cpu=1 --count=1 -timeout 30s github.com/xgxw/toddler-go/internal/services --run TestDemoService
# -v 表示显示测试执行的详细信息
# --run 参数支持正则, 会执行所有以 TestDemoService 开头的单元测试. 

# 运行所有单元测试
go test ./...
```
