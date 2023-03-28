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

type Coinbase struct{}

func NewCoinBaseExchanger() *Coinbase {
	return &Coinbase{}
}

func (ex *Coinbase) TrackCurrencyValue(crypto string) (*models.TrackCurrencyResponse, error) {
	// prepare the url
	url := fmt.Sprintf("%s/exchange-rates?currency=%s", constants.COINBASE_URL, strings.ToUpper(crypto))

	response, err := utils.MakeAPICall(url, http.MethodGet, nil, nil)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var coinbaseResponse models.CoinbaseResponse
	err = json.Unmarshal(body, &coinbaseResponse)
	if err != nil {
		return nil, err
	}
	price, err := strconv.ParseFloat(coinbaseResponse.Data.Rates["USD"], 64)
	if err != nil {
		return nil, err
	}

	return &models.TrackCurrencyResponse{
		Value: price,
	}, nil
}
