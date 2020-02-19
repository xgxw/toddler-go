package toddler

type DemoService interface {
	Check(params map[string]interface{}) (*Result, error)
}

type (
	Result struct {
		OK      bool   `gorm:"column:ok" json:"ok"`
		Limit   uint64 `gorm:"column:limit" json:"limit"`
		Message string `gorm:"column:message" json:"message"`
	}
)

func GetDemoKey() string {
	return "toddler:demo_print"
}
