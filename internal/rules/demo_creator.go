package rules

const DemoRuleName = "demo"

func init() {
	allRuleCreators[DemoRuleName] = NewDemoCreator()
}

type DemoCreator struct{}

var _ IRuleCreator = new(DemoCreator)

func NewDemoCreator() *DemoCreator {
	return &DemoCreator{}
}

func (d *DemoCreator) Create(resource *Resource) IRule {
	return NewDemoRule()
}
