package services

import (
	"github.com/pkg/errors"
	"github.com/xgxw/toddler-go"
	"github.com/xgxw/toddler-go/internal/rules"
)

type RuleService struct {
	ruleFactory *rules.Factory
	allRules    []rules.IRule
}

var _ toddler.RuleService = new(RuleService)

// 注意, 这里要返回外层的Service
func NewRuleService(ruleFactory *rules.Factory) toddler.RuleService {
	allRuleNames := ruleFactory.GetAllRuleNames()
	allRules := make([]rules.IRule, len(allRuleNames))
	for i, name := range allRuleNames {
		allRules[i], _ = ruleFactory.Get(name)
	}

	return &RuleService{
		ruleFactory: ruleFactory,
		allRules:    allRules,
	}
}

func (d *RuleService) Check(params map[string]interface{}) (
	result *toddler.CheckResult, err error) {

	result = &toddler.CheckResult{
		OK:    true,
		Limit: 1 << 63,
	}

	var checkResult *toddler.CheckResult
	for _, rule := range d.allRules {
		checkResult, err = rule.Check(params)
		if err != nil {
			return checkResult, err
		}
		if checkResult == nil {
			return result, errors.Errorf("rule check return nil")
		}
		if !checkResult.OK {
			return checkResult, nil
		}
		if checkResult.Limit < result.Limit {
			result.Limit = checkResult.Limit
			result.Message = checkResult.Message
		}
	}
	return result, nil
}
