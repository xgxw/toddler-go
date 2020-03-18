package cmd

import (
	"log"
	"os"

	"github.com/xgxw/foundation-go/database"
	flog "github.com/xgxw/foundation-go/log"
	"github.com/xgxw/toddler-go"
	"github.com/xgxw/toddler-go/internal/services"
)

type Bootstrap struct {
	Logger  *flog.Logger
	DB      *database.DB
	DemoSvc toddler.DemoService
}

func newBootstrap(opts *Options) *Bootstrap {
	var boot *Bootstrap
	logger := flog.NewLogger(opts.Logging, os.Stdout)

	db, err := database.NewDatabase(opts.DB)
	handleInitError("connect database", err)

	demoSvc := services.NewDemoService()

	boot = &Bootstrap{
		Logger:  logger,
		DB:      db,
		DemoSvc: demoSvc,
	}
	return boot
}

func handleInitError(module string, err error) {
	if err == nil {
		return
	}
	log.Fatalf("init %s failed, err: %s", module, err)
}
