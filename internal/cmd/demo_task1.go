package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var demoTaskCmd = &cobra.Command{
	Use:   "demo_task",
	Short: "demo task",
	Long:  `demo task, 定时任务`,
	Run: func(cmd *cobra.Command, args []string) {
		opts, err := loadOptions()
		handleInitError("load_options", err)
		boot := bootstrap(opts)
		f := boot.RuleFactory
		for _, name := range f.GetAllRuleNames() {
			rule, _ := f.Get(name)
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
