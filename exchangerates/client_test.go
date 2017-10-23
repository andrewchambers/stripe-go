package exchangerates

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestExchangeRatesGet(t *testing.T) {
	// TODO: remove this line once stripe-mock supports /v1/exchange_rates
	t.Skip("not yet supported by stripe-mock")

	rates, err := Get("usd")
	assert.Nil(t, err)
	assert.NotNil(t, rates)
}

func TestExchangeRatesList(t *testing.T) {
	// TODO: remove this line once stripe-mock supports /v1/exchange_rates
	t.Skip("not yet supported by stripe-mock")

	i := List(&stripe.ExchangeRatesListParams{})

	// Verify that we can get at least one exchange_rates
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.ExchangeRates())
}
