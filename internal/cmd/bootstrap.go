package cmd

import (
	"log"
	"os"

	"github.com/xgxw/foundation-go/database"
	flog "github.com/xgxw/foundation-go/log"
	"github.com/xgxw/toddler-go"
	"github.com/xgxw/toddler-go/internal/services"
)

// 文档地址 [doc](/docs/bootstrap.md)

type Bootstrap struct {
	opts     *Options
	logger   *flog.Logger
	db       *database.DB
	demoSvc  toddler.DemoService
	teardown func()
}

func newBootstrap(opts *Options) *Bootstrap {
	return &Bootstrap{
		opts:     opts,
		teardown: func() {},
	}
}

func (b *Bootstrap) Teardown() {
	b.teardown()
}

func (b *Bootstrap) GetLogger() *flog.Logger {
	if b.logger == nil {
		b.logger = flog.NewLogger(b.opts.Logging, os.Stdout)
	}
	return b.logger
}

func (b *Bootstrap) GetDB() *database.DB {
	if b.db == nil {
		db, err := database.NewDatabase(b.opts.DB)
		handleInitError("connect database", err)
		b.addTeardown(func() { db.Close() })
	}
	return b.db
}

func (b *Bootstrap) GetDemoSvc() toddler.DemoService {
	if b.demoSvc == nil {
		b.demoSvc = services.NewDemoService()
	}
	return b.demoSvc
}

func (b *Bootstrap) addTeardown(newTeardown func()) {
	teardown := b.teardown
	b.teardown = func() {
		teardown()
		newTeardown()
	}
}

func handleInitError(module string, err error) {
	if err == nil {
		return
	}
	log.Fatalf("init %s failed, err: %s", module, err)
}
