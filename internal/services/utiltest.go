package services

import (
	"time"

	"github.com/xgxw/foundation-go/database"
	"github.com/xgxw/toddler-go"
)

/*
	UtilTestService 用于添加单元测试的使用示例. 所以不在外层定义该服务了,
	服务和结构体都定义在文件内了.
	依赖 DemoService 是为了添加 gomock 的使用示例, 所以首先需要执行 make mock
	生成 DemoService 的mock文件. 具体参考 Makefile#mock
*/

type UtilTestService struct {
	db      *database.DB
	demoSvc toddler.DemoService
}

type UtilTestStruct struct {
	ID   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`

	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}

func (UtilTestStruct) TableName() string {
	return "util_test_structs"
}

// 注意, 这里要返回外层的Service
func NewUtilTestService(db *database.DB, demoSvc toddler.DemoService) *UtilTestService {
	return &UtilTestService{
		db:      db,
		demoSvc: demoSvc,
	}
}

// DoSomething 测试调用其他服务
func (dSvc *UtilTestService) DoSomething() (string, error) {
	request := new(toddler.Request)
	response, err := dSvc.demoSvc.DoSomething(request)
	if err != nil {
		return "", err
	}
	return response.Msg, nil
}

// Create 创建数据库记录示例
func (dSvc *UtilTestService) Create(name string) (*UtilTestStruct, error) {
	utStruct := &UtilTestStruct{
		Name: name,
	}
	return utStruct, dSvc.db.Create(utStruct).Error
}

func (dSvc *UtilTestService) Update(id int, name string) error {
	return dSvc.db.Model(&UtilTestStruct{}).Where("id=?", id).Update("name", name).Error
}

func (dSvc *UtilTestService) Get(name string) (*UtilTestStruct, error) {
	utStruct := new(UtilTestStruct)
	err := dSvc.db.Where("name=?", name).First(utStruct).Error
	return utStruct, err
}

func (dSvc *UtilTestService) Delete(name string) error {
	return dSvc.db.Where("name=?", name).Delete(&UtilTestStruct{}).Error
}
