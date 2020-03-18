package controllers

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	flog "github.com/xgxw/foundation-go/log"
	"github.com/xgxw/toddler-go"
	"github.com/xgxw/toddler-go/internal/codes"
	"github.com/xgxw/toddler-go/pb"
)

// DemoController is GRPC controller
type DemoController struct {
	logger  *flog.Logger
	demoSvc toddler.DemoService
}

var _ pb.DemoServiceServer = &DemoController{}

func NewDemoController(logger *flog.Logger, demoSvc toddler.DemoService) *DemoController {
	return &DemoController{
		logger:  logger,
		demoSvc: demoSvc,
	}
}

func (demoCtl *DemoController) DoSomething(ctx context.Context,
	r *pb.Request) (resp *pb.Response, err error) {
	if r.Validate() != nil {
		return &pb.Response{}, codes.BadRequestError
	}
	request := &toddler.Request{
		ID:   int(r.GetId()),
		Name: r.GetName(),
	}
	response, err := demoCtl.demoSvc.DoSomething(request)
	if err != nil {
		// error记录: WithField 用于表示字段, 存储后便于检索.
		// WithError 封装error信息
		// Error 添加错误描述
		demoCtl.logger.WithError(err).
			WithField("func", "doSomething").
			Error("xxx error")
		return &pb.Response{}, err
	}
	return &pb.Response{
		Ok:  response.OK,
		Msg: response.Msg,
	}, nil
}

func (demoCtl *DemoController) DoSomethingHttp(ctx echo.Context) (err error) {
	request := new(toddler.Request)
	err = ctx.Bind(request)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	response, err := demoCtl.demoSvc.DoSomething(request)
	if err != nil {
		demoCtl.logger.WithError(err).
			WithField("func", "doSomething").
			Error("xxx error")
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, response)
}
