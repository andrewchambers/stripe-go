package stripe

// ExchangeRates is the resource representing the currency exchange rates at
// a given time.
type ExchangeRates struct {
	ID    string               `json:"id"`
	Rates map[Currency]float64 `json:"rates"`
}

// ExchangeRatesList is a list of exchange rates as retrieved from a list endpoint.
type ExchangeRatesList struct {
	ListMeta
	Values []*ExchangeRates `json:"data"`
}

// ExchangeRatesListParams are the parameters allowed during ExchangeRates listing.
type ExchangeRatesListParams struct {
	ListParams `form:"*"`
}
