package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cobra"
	"github.com/xgxw/toddler-go/internal/controllers"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "http server",
	Long:  `服务端`,
	Run: func(cmd *cobra.Command, args []string) {
		opts, err := loadOptions()

		boot := newBootstrap(opts)
		defer boot.Teardown()

		logger := boot.GetLogger()
		demoCtl := controllers.NewDemoController(logger, boot.GetDemoSvc())

		e := echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     opts.Server.CorsAllowOrigin,
			AllowCredentials: true,
		}))

		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "enjoy yourself!")
		})

		v1 := e.Group("/v1")

		{
			demo := v1.Group("demo")
			demo.GET("/", demoCtl.DoSomethingHttp)
		}

		quit := make(chan os.Signal, 1)
		go func() {
			// 当程序较多/HTTP设置较多时, 可以单独封装Server组件, 在组件内计算这些值
			address := fmt.Sprintf("%s:%d", opts.Server.Host, opts.Server.Port)
			err = e.Start(address)
			if err != nil {
				logger.Fatal("start echo error, error is ", err)
				quit <- os.Interrupt
			}
		}()
		signal.Notify(quit, os.Interrupt)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

type (
	ServerOption struct {
		Host            string   `mapstructure:"host" yaml:"host"`
		Port            uint     `mapstructure:"port" yaml:"port"`
		CorsAllowOrigin []string `mapstructure:"cors_allow_origin"`
		DefaultExpired  int64    `mapstructure:"default_expired"`
	}
)
