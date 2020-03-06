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

// DemoController is Agent GRPC controller
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

func (demoCtl *DemoController) Check(ctx context.Context,
	r *pb.CheckRequest) (resp *pb.CheckResponse, err error) {
	if r.Validate() != nil {
		return &pb.CheckResponse{}, codes.BadRequestError
	}
	var m = make(map[string]interface{})
	checkResult, err := demoCtl.demoSvc.Check(m)
	if err != nil {
		demoCtl.logger.WithError(err).
			Error("check error")
		return &pb.CheckResponse{}, err
	}
	return &pb.CheckResponse{
		Ok:    checkResult.OK,
		Limit: checkResult.Limit,
		Msg:   checkResult.Message,
	}, nil
}

func (demoCtl *DemoController) CheckHTTP(ctx echo.Context) (err error) {
	return ctx.String(http.StatusOK, "this is controller")
}
