//go:build integrationtestderivativecontract

package integrationtestderivativecontract

import (
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestDerivativesOrderBook(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Derivative().UnifiedMargin().DerivativesOrderBook(bybit.DerivativesOrderBookParam{
		Symbol:   bybit.SymbolDerivativeBTCUSDT,
		Category: bybit.CategoryDerivativeLinear,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/derivatives-public-order-book-l2.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
