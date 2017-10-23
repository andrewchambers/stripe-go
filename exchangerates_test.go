package stripe

import (
	"encoding/json"
	"testing"
)

func TestExchangeRatesUnmarshal(t *testing.T) {
	exchangeRatesData := map[string]interface{}{
		"id":     "usd",
		"object": "exchange_rates",
		"rates": map[string]interface{}{
			"eur": 0.845876,
		},
	}

	bytes, err := json.Marshal(&exchangeRatesData)
	if err != nil {
		t.Error(err)
	}

	var exchangeRates ExchangeRates
	err = json.Unmarshal(bytes, &exchangeRates)
	if err != nil {
		t.Error(err)
	}

	if exchangeRates.Rates == nil {
		t.Errorf("Problem deserializing rates, got nothing.")
	}

	if exchangeRates.Rates["eur"] != 0.845876 {
		t.Errorf("Problem deserializing rates[eur], got %v", exchangeRates.Rates["eur"])
	}
}
