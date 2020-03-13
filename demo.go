package toddler

import "time"

type DemoService interface {
	Check(params map[string]interface{}) (*Result, error)
}

type (
	Demo struct {
		ID   int    `gorm:"column:id" json:"id"`
		Name string `gorm:"column:name" json:"name"`

		CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
		UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
		DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
	}
	Result struct {
		OK      bool   `gorm:"column:ok" json:"ok"`
		Limit   uint64 `gorm:"column:limit" json:"limit"`
		Message string `gorm:"column:message" json:"message"`
	}
)

func (Demo) TableName() string {
	return "demos"
}

func GetDemoKey() string {
	return "toddler:demo_print"
}
