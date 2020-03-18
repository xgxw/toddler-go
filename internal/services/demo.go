package services

import (
	"github.com/xgxw/toddler-go"
)

// DemoService 被确立为基础服务, 被其他服务调用, 后续改动要慎重操作

type DemoService struct{}

// 注意, 这里要返回外层的Service
func NewDemoService() toddler.DemoService {
	return &DemoService{}
}

var _ toddler.DemoService = new(DemoService)

// Create 创建数据库记录示例
func (dSvc *DemoService) DoSomething(request *toddler.Request) (*toddler.Response, error) {
	return &toddler.Response{}, nil
}
