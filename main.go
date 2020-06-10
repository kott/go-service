package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	rcontext "github.com/kott/go-service/context"
	"github.com/kott/go-service/log"
	"github.com/kott/go-service/middleware"
)

const (
	servicesProfile = "SERVICES_PROFILE"
	dockerProfile   = "docker"
)

func init() {
	profile := os.Getenv(servicesProfile)
	if profile == dockerProfile { // all env vars are set in the container already
		viper.AutomaticEnv()
		viper.AllowEmptyEnv(true)
	} else {
		viper.SetConfigName(strings.ToLower(profile))
		viper.SetConfigType("env")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("an error occurred reading the config file: %s ", err))
		}
	}
}

func main() {
	app := gin.New()

	app.Use(middleware.PersistContext())
	app.Use(middleware.RequestLogger())
	app.Use(middleware.ForceJSON())
	app.Use(middleware.Recover())

	app.NoRoute(middleware.NoRoute())
	app.NoMethod(middleware.NoMethod())

	app.GET("/", func(c *gin.Context) {
		ctx := rcontext.GetReqCtx(c)
		log.Info(ctx, "Info log")
		c.JSON(http.StatusOK, struct{}{})
	})

	if err := app.Run(fmt.Sprintf("%s:%s", viper.Get("host"), viper.Get("port"))); err != nil {
		log.Fatal(context.Background(), err.Error())
	}
}
