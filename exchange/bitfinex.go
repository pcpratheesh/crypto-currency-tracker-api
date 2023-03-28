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

type BitFinex struct{}

func NewBitFinexExchanger() *BitFinex {
	return &BitFinex{}
}

func (ex *BitFinex) TrackCurrencyValue(from string, to string) (*models.TrackCurrencyResponse, error) {
	// prepare the url
	url := fmt.Sprintf("%s/ticker/t%s%s", constants.BITFINEX_URL, strings.ToUpper(from), strings.ToUpper(to))

	response, err := utils.MakeAPICall(url, http.MethodGet, nil, nil)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var bitfinexResponse []interface{}
	err = json.Unmarshal(body, &bitfinexResponse)
	if err != nil {
		return nil, err
	}
	fmt.Println(bitfinexResponse[6])

	return &models.TrackCurrencyResponse{
		From:  from,
		To:    to,
		Value: bitfinexResponse[6],
	}, nil

}
