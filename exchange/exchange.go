package exchange

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const baseURL = "https://openexchangerates.org/api/latest.json"
const appID = "34a3932f56234f4796598c4cd5ac8969"

type exchangeRates struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

func GetCurrentEchangeRate(currency string) float64 {
	url := fmt.Sprintf("%s?app_id=%s", baseURL, appID)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Couldn't get a response: %v", err)
	}

	var data exchangeRates

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Couldn't decode response body: %v", err)
	}

	rate, ok := data.Rates[currency]
	if !ok {
		log.Fatalf("CHF rate not found")
	}

	return rate
}
