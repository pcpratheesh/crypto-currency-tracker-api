package exchange

import (
	"fmt"
	"strings"

	"github.com/pcpratheesh/crypto-currency-tracker-api/constants"
	"github.com/pcpratheesh/crypto-currency-tracker-api/models"
)

type ExchangerInterface interface {
	TrackCurrencyValue(string) (*models.TrackCurrencyResponse, error)
}

func ChooseExchanger(ex string) (ExchangerInterface, error) {
	switch strings.ToUpper(ex) {
	case constants.EXCHANGE_BINANCE:
		return NewBinanceExchanger(), nil
	case constants.EXCHANGE_COINBASE:
		return NewCoinBaseExchanger(), nil
	case constants.EXCHANGE_BITFINEX:
		return NewBitFinexExchanger(), nil
	}

	return nil, fmt.Errorf("%v exchanger not borded yet", ex)
}
