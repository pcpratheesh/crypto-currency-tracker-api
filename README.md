# crypto-currency-tracker-api
A RESTful API service that tracks the prices and convert of various cryptocurrencies using real-time data from different cryptocurrency exchanges.


## Installation
- Clone the repository: git clone https://github.com/pcpratheesh/crypto-currency-tracker-api.git
- Navigate to the project directory: cd crypto-currency-tracker-api
- Install the dependencies: go get ./...

## Usage
To use the API service, simply run the following command in the project directory:

    go run main.go

The API service will start listening on port 8085 by default. You can change the port by setting the PORT environment variable.


## Endpoints
The following endpoints are available:

`GET /api/v1/convert/:exchange`
Retrieves the current prices of various cryptocurrencies from different exchanges.

`Json Body`
```json
    {
        "crypto":"BTC",
        "base":"USD"
    }
```
## Sample API Request

- With bitfinex
```sh
    curl --request POST \
    --url http://localhost:8085/api/v1/track/bitfinex \
    --header 'Content-Type: application/json' \
    --data '{
        "crypto":"BTC",
        "base":"USD"
    }'
```

- With binance
```sh
    curl --request POST \
    --url http://localhost:8085/api/v1/track/binance \
    --header 'Content-Type: application/json' \
    --data '{
        "crypto":"BTC",
        "base":"USDT"
    }'
```

- With coinbase
```sh
    curl --request POST \
    --url http://localhost:8085/api/v1/track/coinbase \
    --header 'Content-Type: application/json' \
    --data '{
        "crypto":"BTC",
        "base":"USDT"
    }'
```

## Supported Exchanges
The following cryptocurrency exchanges are currently supported by the application:
- Coinbase
- Binance
- Bitfinex


## Contributing

If you find a bug or would like to contribute to the development of the application, please submit a pull request or create an issue on the GitHub repository.

## License
This project is licensed under the MIT License - see the LICENSE.md file for details.