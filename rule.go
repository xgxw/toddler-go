package toddler

type RuleService interface {
	Check(params map[string]interface{}) (*CheckResult, error)
}

type (
	CheckResult struct {
		OK      bool   `gorm:"column:ok" json:"ok"`
		Limit   uint64 `gorm:"column:limit" json:"limit"`
		Message string `gorm:"column:message" json:"message"`
	}
)
