package rules

import "github.com/xgxw/toddler-go"

// IRule rule interface
type IRule interface {
	// Check 检查是否符合规则
	Check(map[string]interface{}) (*toddler.CheckResult, error)
	MakeEffective() error
}
