package rules

import (
	"fmt"

	"github.com/xgxw/foundation-go/utils"
	"github.com/xgxw/toddler-go"
)

const DemoLimit = uint64(30000)

var (
	DemoCheckSuccess = &toddler.CheckResult{
		OK:    true,
		Limit: DemoLimit,
	}
	DemoCheckFail = &toddler.CheckResult{
		OK:      false,
		Limit:   DemoLimit,
		Message: "xxxx",
	}
)

type DemoRule struct{}

var _ IRule = new(DemoRule)

func NewDemoRule() *DemoRule {
	return &DemoRule{}
}

type DemoCheckRequest struct {
	Amount uint64 `gorm:"column:amount" json:"amount"`
}

func (d *DemoRule) Check(params map[string]interface{}) (*toddler.CheckResult, error) {
	_, err := d.parseCheckParams(params)
	if err != nil {
		return DemoCheckFail, err
	}

	// 假装这里有事做
	fmt.Println("demo rule Checked")

	return DemoCheckSuccess, nil
}

func (d *DemoRule) parseCheckParams(
	params map[string]interface{}) (*DemoCheckRequest, error) {
	request := &DemoCheckRequest{}
	err := utils.FillStruct(request, params)
	return request, err
}

func (d *DemoRule) MakeEffective() error {
	fmt.Println("demo rule effective")
	// 假装这里有事做
	return nil
}
