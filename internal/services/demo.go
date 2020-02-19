package services

import (
	"github.com/pkg/errors"
	"github.com/xgxw/toddler-go"
	"github.com/xgxw/toddler-go/internal/rules"
)

type DemoService struct {
	ruleFactory *rules.Factory
	allRules    []rules.IRule
}

func NewDemoService(ruleFactory *rules.Factory) *DemoService {

	allRuleNames := ruleFactory.GetAllRuleNames()
	allRules := make([]rules.IRule, len(allRuleNames))
	for i, name := range allRuleNames {
		allRules[i], _ = ruleFactory.Get(name)
	}

	return &DemoService{
		ruleFactory: ruleFactory,
		allRules:    allRules,
	}
}

var _ toddler.DemoService = new(DemoService)

func (d *DemoService) Check(params map[string]interface{}) (
	result *toddler.Result, err error) {

	result = &toddler.Result{
		OK:    true,
		Limit: 1 << 63,
	}

	var checkResult *toddler.Result
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
