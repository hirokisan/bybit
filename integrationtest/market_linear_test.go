//go:build integrationtest

package integrationtest

import (
	"testing"

	"github.com/hirokisan/bybit"
	"github.com/hirokisan/bybit/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestLinearTickers(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Market().LinearTickers(bybit.SymbolUSDTBTC)
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-tickers.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
