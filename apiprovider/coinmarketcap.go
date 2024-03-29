package apiprovider

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// CoinMarketCapURL is url of CoinMarketCap
const CoinMarketCapURL = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest"

// CoinMarketCap will set the basic data of CoinMarketCap needs
type CoinMarketCap struct {
	URL    string
	APIKey string
}

// Quote represent cryptocurreny's value of a specific currency
type Quote struct {
	Price     float64
	Volume24h float64 `json:"volume_24h"`
}

// Data represent a cryptocurrency status
type Data struct {
	ID          int
	Name        string
	Quote       map[string]Quote
	LastUpdated time.Time `json:"last_updated"`
}

// Status represent the status of response
type Status struct {
	Timestamp   time.Time
	Elapsed     int
	CreditCount int `json:"credit_count"`
}

// Response is CoinMarketCap's quotes reponse
type Response struct {
	Data   map[string]Data
	Status Status
}

// GetLatestPrice will get latest BTC price with USD
func (c *CoinMarketCap) GetLatestPrice(currency Currency) (float64, error) {
	currencyID := map[Currency]int{
		Usd: 2781,
		Twd: 2811,
	}[currency]

	request, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return 0, err
	}

	q := url.Values{}
	q.Add("id", "1")
	q.Add("convert_id", strconv.Itoa(currencyID))

	request.Header.Set("Accepts", "application/json")
	request.Header.Add("X-CMC_PRO_API_KEY", c.APIKey)
	request.URL.RawQuery = q.Encode()

	client := &http.Client{}
	r, err := client.Do(request)
	if err != nil {
		return 0, err
	}

	var response Response
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return 0, err
	}

	return response.Data["1"].Quote[strconv.Itoa(currencyID)].Price, nil
}
