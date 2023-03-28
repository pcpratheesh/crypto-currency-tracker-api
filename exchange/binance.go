package exchange

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/pcpratheesh/crypto-currency-tracker-api/constants"
	"github.com/pcpratheesh/crypto-currency-tracker-api/models"
	"github.com/pcpratheesh/crypto-currency-tracker-api/utils"
)

type Binance struct{}

func NewBinanceExchanger() *Binance {
	return &Binance{}
}

func (ex *Binance) TrackCurrencyValue(crypto string) (*models.TrackCurrencyResponse, error) {
	// prepare the url
	url := fmt.Sprintf("%s/ticker/price?symbol=%s%s", constants.BINANCE_URL, strings.ToUpper(crypto), strings.ToUpper("USDT"))

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

	price, err := strconv.ParseFloat(binanceTickerResponse.Price, 64)
	if err != nil {
		return nil, err
	}

	return &models.TrackCurrencyResponse{
		Value: price,
	}, nil
}
