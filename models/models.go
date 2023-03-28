package models

type BinanceTickerResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type CoinbaseResponse struct {
	Data struct {
		Base     string `json:"base"`
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"data"`
}

type TrackCurrencyRequest struct {
	Crypto string `json:"crypto" validate:"required"`
	Base   string `json:"base" validate:"required"`
}

type TrackCurrencyResponse struct {
	From  string      `json:"from"`
	To    string      `json:"to"`
	Value interface{} `json:"value"`
}
