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

type Price struct {
	Symbol string      `json:"symbol"`
	Price  interface{} `json:"price"`
	Error  error       `json:"error"`
}

func (c *Controller) TrackCurrencyValue(ctx echo.Context) error {
	// read the request
	var request = models.TrackCurrencyRequest{}
	var exc = ctx.Param("exchange")

	if err := ctx.Bind(&request); err != nil {
		return utils.HTTPErrorHandler(ctx, err, http.StatusBadRequest)
	}

	prices := make(chan Price, len(request.Currencies))

	// choose the exchanger
	// choose the exchanger
	manager, err := exchange.ChooseExchanger(exc)
	if err != nil {
		return utils.HTTPErrorHandler(ctx, err, http.StatusBadRequest)
	}

	for _, currency := range request.Currencies {
		go func(currency string) {
			// trigger api
			data, err := manager.TrackCurrencyValue(currency)
			if err != nil {
				prices <- Price{
					Error: utils.HTTPErrorHandler(ctx, err, http.StatusBadRequest),
				}
				return
			}

			prices <- Price{
				Symbol: currency,
				Price:  data.Value,
			}

		}(currency)
	}

	result := make(map[string]interface{})
	for range request.Currencies {
		price := <-prices

		if price.Error != nil {
			logrus.Error(price.Error)
			continue
		}

		result[price.Symbol] = price.Price
	}

	//respond back
	return ctx.JSON(http.StatusOK, result)
}
