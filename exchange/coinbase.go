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

type Coinbase struct{}

func NewCoinBaseExchanger() *Coinbase {
	return &Coinbase{}
}

func (ex *Coinbase) TrackCurrencyValue(from string, to string) (*models.TrackCurrencyResponse, error) {
	// prepare the url
	url := fmt.Sprintf("%s/prices/%s-%s/spot", constants.COINBASE_URL, strings.ToUpper(from), strings.ToUpper(to))

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

	return &models.TrackCurrencyResponse{
		From:  from,
		To:    to,
		Value: coinbaseResponse.Data.Amount,
	}, nil
}
