package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pcpratheesh/crypto-currency-tracker-api/controller"
	"github.com/pcpratheesh/crypto-currency-tracker-api/utils"
	"github.com/sirupsen/logrus"
)

var (
	PORT = 8085
)

func Run() {
	// initializing the controller
	controller := controller.NewController()

	router := echo.New()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.Validator = utils.RegisterValidator()

	// base router
	router.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "server is up")
	})

	v1 := router.Group("/api/v1")
	// track route
	v1.POST("/track/:exchange", controller.TrackCurrencyValue)

	data, err := json.MarshalIndent(router.Routes(), "", "  ")
	if err != nil {
		panic(err)
	}

	ll := logrus.New()
	ll.Info(string(data))
	ll.Info("server is running....")

	// Start server
	router.Logger.Fatal(router.Start(fmt.Sprintf(":%d", PORT)))
}
