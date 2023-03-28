package exchange

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pcpratheesh/crypto-currency-tracker-api/constants"
	"github.com/pcpratheesh/crypto-currency-tracker-api/models"
	"github.com/pcpratheesh/crypto-currency-tracker-api/utils"
)

type Binance struct{}

func NewBinanceExchanger() *Binance {
	return &Binance{}
}

func (ex *Binance) TrackCurrencyValue(from string, to string) (*models.TrackCurrencyResponse, error) {
	// prepare the url
	url := fmt.Sprintf("%s/ticker/price?symbol=%s%s", constants.BINANCE_URL, strings.ToUpper(from), strings.ToUpper(to))

	response, err := utils.MakeAPICall(url, http.MethodGet, nil, nil)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var binanceTickerResponse models.BinanceTickerResponse
	err = json.Unmarshal(body, &binanceTickerResponse)
	if err != nil {
		return nil, err
	}

	return &models.TrackCurrencyResponse{
		From:  from,
		To:    to,
		Value: binanceTickerResponse.Price,
	}, nil
}
