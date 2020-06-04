package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	setLogLevel(e.Logger, log.INFO)
	e.GET("/", func(c echo.Context) error {
		c.Logger().Info("")
		return c.JSONPretty(http.StatusOK, struct{Msg string}{Msg: "hello"},"\t")
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func setLogLevel(logger echo.Logger, lvl log.Lvl) {
	if l, ok := logger.(*log.Logger); ok {
		l.SetLevel(lvl)
	}
}
