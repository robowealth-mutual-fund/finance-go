package crypto

import (
	"testing"

	finance "github.com/robowealth-mutual-fund/finance-go"
	tests "github.com/robowealth-mutual-fund/finance-go/testing"
	"github.com/stretchr/testify/assert"
)

func TestGetCryptoPair(t *testing.T) {
	tests.SetMarket(finance.MarketStateRegular)

	q, err := Get(tests.TestCryptoPairSymbol)

	assert.Nil(t, err)
	assert.NotNil(t, q)
	assert.Equal(t, finance.MarketStateRegular, q.MarketState)
	assert.Equal(t, tests.TestCryptoPairSymbol, q.Symbol)
}
