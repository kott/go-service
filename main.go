package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	rcontext "github.com/kott/go-service/context"
	"github.com/kott/go-service/log"
	"github.com/kott/go-service/middleware"
)

func main() {

	app := gin.New()

	app.Use(middleware.PersistContext())
	app.Use(middleware.RequestLogger())

	app.GET("/", func(c *gin.Context) {
		ctx := rcontext.GetReqCtx(c)
		log.Info(ctx, "Info log")
		c.JSON(http.StatusOK, struct{}{})
	})

	if err := app.Run(fmt.Sprintf("%s:%s", "0.0.0.0", "8080")); err != nil {
		log.Fatal(context.Background(),"THINGS WENT HORRIBLY WRONG")
	}
}
