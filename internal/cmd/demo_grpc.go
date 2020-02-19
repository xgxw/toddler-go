package cmd

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/rs/xid"
	"github.com/spf13/cobra"
	"github.com/xgxw/toddler-go/internal/controllers"
	"github.com/xgxw/toddler-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var demoGRPCCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo",
	Long:  `demo`,
	Run: func(cmd *cobra.Command, args []string) {
		opts, err := loadOptions()
		handleInitError("load_options", err)
		boot := bootstrap(opts)

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", opts.Demo.Port))
		handleInitError("grpc_listen", err)

		demoCtl := controllers.NewDemoController(boot.Logger, boot.DemoSvc)

		logger := boot.Logger.WithField("scope", "demo")
		gs := grpc.NewServer(
			grpc.KeepaliveParams(keepalive.ServerParameters{
				Time: 5 * time.Second,
			}),
			grpc_middleware.WithUnaryServerChain(
				grpc_ctxtags.UnaryServerInterceptor(
					grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor),
					grpc_ctxtags.WithFieldExtractor(func(fullMethod string, req interface{}) map[string]interface{} {
						fields := map[string]interface{}{"request_id": xid.New().String()}
						return fields
					}),
				),
				grpc_logrus.UnaryServerInterceptor(logger),
				grpc_logrus.PayloadUnaryServerInterceptor(logger, func(ctx context.Context, fullMethodName string, servingObject interface{}) bool { return true }),
				grpc_validator.UnaryServerInterceptor(),
			),
		)
		pb.RegisterDemoServiceServer(gs, demoCtl)

		quit := make(chan os.Signal, 1)
		go func() {
			boot.Logger.Infof("grpc server start at port %d...", opts.Demo.Port)
			err = gs.Serve(lis)
			if err != nil {
				boot.Logger.Fatalf("start server error, error is %v ", err)
				quit <- os.Interrupt
			}
		}()
		signal.Notify(quit, os.Interrupt)
		<-quit

		gs.GracefulStop()
	},
}

type DemoOption struct {
	Port int `mapstructure:"port" yaml:"port"`
}

func init() {
	rootCmd.AddCommand(demoGRPCCmd)
}
