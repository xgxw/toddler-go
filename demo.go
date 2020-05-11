package toddler

import "time"

// 使用 interface 的好处参考 link:/docs/services.adoc

type DemoService interface {
	DoSomething(*Request) (*Response, error)
}

type (
	Request struct {
		ID   int    `gorm:"column:id" json:"id"`
		Name string `gorm:"column:name" json:"name"`
	}
	Response struct {
		OK  bool   `gorm:"column:ok" json:"ok"`
		Msg string `gorm:"column:msg" json:"msg"`
	}
)

type (
	// DemoStructDTO 是 DemoStruct 对外展示的结构体,
	// DemoStruct 即 DemoStructDO, 是数据库存储的结构体.
	// 具体参考 link:/docs/namings[命名规范]
	DemoStructDTO DemoStruct
	// DemoStruct 结构体定义示例, gorm 会自动设置 CreatedAt/UpdatedAt/DeletedAt
	DemoStruct struct {
		ID   int    `gorm:"column:id" json:"id"`
		Name string `gorm:"column:name" json:"name"`

		CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
		UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
		DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
	}
)

// TableName gorm 规范, 需要添加 TableName 以便gorm获取表名.
// 规范为 "业务名+s", 连字符模式. 如 user=> users.
func (DemoStruct) TableName() string {
	return "demo_structs"
}
