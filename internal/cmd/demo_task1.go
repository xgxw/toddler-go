package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xgxw/toddler-go/internal/rules"
)

var demoTaskCmd = &cobra.Command{
	Use:   "demo_task",
	Short: "demo task",
	Long:  `demo task, 定时任务`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		//opts, err := loadOptions()
		//handleInitError("load_options", err)
		resource := &rules.Resource{}
		ruleFactory := rules.NewFactory(resource)

		for _, name := range ruleFactory.GetAllRuleNames() {
			rule, _ := ruleFactory.Get(name)
			err = rule.MakeEffective()
			if err != nil {
				log.Fatalf("error in exec rule make effective. error: %v", err)
			}
		}
		log.Println("done")
	},
}

func init() {
	rootCmd.AddCommand(demoTaskCmd)
}
