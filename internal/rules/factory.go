package rules

import (
	"github.com/pkg/errors"
)

type Factory struct {
	resource *Resource
}

func NewFactory(resource *Resource) *Factory {
	return &Factory{
		resource: resource,
	}
}

// allRuleCreators 存储所有rule的Creator. 在每个rule文件init函数中设置
var allRuleCreators = make(map[string]IRuleCreator, 0)

// Resource 统一资源管理, 存储所有 rule 需要的资源.
// factory 将 Resource 传入 creator, creator 取其中资源, 交于NewXXRule.
type Resource struct{}

// IRuleCreator 为工厂提供统一的创建方法
type IRuleCreator interface {
	Create(resource *Resource) IRule
}

// Get 返回指定的 Rule, 如果找不到, 会返回error
func (f *Factory) Get(name string) (IRule, error) {
	ruleCreator, ok := allRuleCreators[name]
	if !ok {
		return nil, errors.New("rule not found. rule name: " + name)
	}

	return ruleCreator.Create(f.resource), nil
}

func (f *Factory) GetAllRuleNames() []string {
	all := f.GetAllRuleCreators()
	var names = make([]string, len(all))
	i := 0
	for k, _ := range all {
		names[i] = k
		i++
	}
	return names
}

func (f *Factory) GetAllRuleCreators() map[string]IRuleCreator {
	return allRuleCreators
}

/*
	1. 工厂只负责生产, 不负责单例化
*/
