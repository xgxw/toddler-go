package helpers

import (
	"github.com/xgxw/foundation-go/database"
	"github.com/xgxw/toddler-go"
)

type DemoHelper struct {
	db *database.DB
}

// 正常应该返回 toddler.CURDService, 但是此处为了方便, 没有定义外层Service
func NewDemoHelper(db *database.DB) *DemoHelper {
	return &DemoHelper{
		db: db,
	}
}

// Create 创建数据库记录示例
func (helper *DemoHelper) Create(name string) (*toddler.Demo, error) {
	demo := &toddler.Demo{
		Name: name,
	}
	return demo, helper.db.Create(demo).Error
}

func (helper *DemoHelper) Update(id int, name string) error {
	return helper.db.Model(&toddler.Demo{}).Where("id=?", id).Update("name", name).Error
}

func (helper *DemoHelper) Get(name string) (*toddler.Demo, error) {
	demo := new(toddler.Demo)
	err := helper.db.Where("name=?", name).First(demo).Error
	return demo, err
}

func (helper *DemoHelper) Delete(name string) error {
	return helper.db.Where("name=?", name).Delete(&toddler.Demo{}).Error
}
