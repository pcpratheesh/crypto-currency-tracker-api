package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pcpratheesh/crypto-currency-tracker-api/exchange"
	"github.com/pcpratheesh/crypto-currency-tracker-api/models"
	"github.com/pcpratheesh/crypto-currency-tracker-api/utils"
	"github.com/sirupsen/logrus"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) TrackCurrencyValue(ctx echo.Context) error {
	// read the request
	var request = models.TrackCurrencyRequest{}
	var exc = ctx.Param("exchange")

	if err := ctx.Bind(&request); err != nil {
		return utils.HTTPErrorHandler(ctx, err, http.StatusBadRequest)
	}

	// choose the exchanger
	manager, err := exchange.ChooseExchanger(exc)
	if err != nil {
		return utils.HTTPErrorHandler(ctx, err, http.StatusBadRequest)
	}
	// trigger api
	data, err := manager.TrackCurrencyValue(request.Crypto, request.Base)
	if err != nil {
		logrus.Errorf("[TrackCurrencyValue] : %v", err)
		return utils.HTTPErrorHandler(ctx, err, http.StatusBadRequest)
	}
	//respond back
	return ctx.JSON(http.StatusOK, data)
}
