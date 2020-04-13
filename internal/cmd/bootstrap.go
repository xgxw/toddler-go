package cmd

import (
	"log"
	"os"

	"github.com/xgxw/foundation-go/database"
	flog "github.com/xgxw/foundation-go/log"
	"github.com/xgxw/toddler-go"
	"github.com/xgxw/toddler-go/internal/services"
)

/*
	Bootstrap 提供统一的资源创建. 这么做有如下优点
	1. 依赖注入/控制反转: 将资源的创建操作统一放到 Bootstrap, 由 Bootstrap 协调.
	2. 门面模式: 提供一个统一的类去访问多个子系统的多个不同的资源.

	这么做的好处有
	1. 后续资源创建方法更改后, 只需要更改Bootstrap中的创建方法, 不需要改动其他地方.
	2. 统一的创建减少代码重复度, 让代码更简洁直观.

	刚开始时, 是以成员字段的方式提供资源, 现在改为 GetLogger() 的方式提供资源.
	新的方式更灵活, 更能适应多个Command的情况.
	如 server/grpc 需要 logger/db/cache/demoSvc
	task1 只需要 db
	task2 只需要 demoSvc

	使用这种方式可以更灵活的初始化, 避免task1运行时添加 cache/demoSvc 的配置
*/

type Bootstrap struct {
	opts    *Options
	logger  *flog.Logger
	db      *database.DB
	demoSvc toddler.DemoService
}

func newBootstrap(opts *Options) *Bootstrap {
	return &Bootstrap{
		opts: opts,
	}
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
		b.db = db
	}
	return b.db
}

func (b *Bootstrap) GetDemoSvc() toddler.DemoService {
	if b.demoSvc == nil {
		b.demoSvc = services.NewDemoService()
	}
	return b.demoSvc
}

func handleInitError(module string, err error) {
	if err == nil {
		return
	}
	log.Fatalf("init %s failed, err: %s", module, err)
}
