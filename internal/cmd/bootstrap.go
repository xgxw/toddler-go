package cmd

import (
	"log"
	"os"

	"github.com/xgxw/foundation-go/database"
	flog "github.com/xgxw/foundation-go/log"
	"github.com/xgxw/toddler-go"
	"github.com/xgxw/toddler-go/internal/rules"
	"github.com/xgxw/toddler-go/internal/services"
)

type Bootstrap struct {
	Logger      *flog.Logger
	DB          *database.DB
	RuleFactory *rules.Factory
	DemoSvc     toddler.DemoService
}

// bootstrap 用于注入依赖
func newBootstrap(opts *Options) *Bootstrap {
	var boot *Bootstrap
	logger := flog.NewLogger(opts.Logging, os.Stdout)

	db, err := database.NewDatabase(opts.DB)
	handleInitError("connect database", err)

	resource := &rules.Resource{}

	ruleFactory := rules.NewFactory(resource)
	demoSvc := services.NewDemoService(ruleFactory)

	boot = &Bootstrap{
		Logger:      logger,
		DB:          db,
		RuleFactory: ruleFactory,
		DemoSvc:     demoSvc,
	}
	return boot
}

func handleInitError(module string, err error) {
	if err == nil {
		return
	}
	log.Fatalf("init %s failed, err: %s", module, err)
}
