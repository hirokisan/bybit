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
	res, err := client.Future().Common.LinearTickers(bybit.SymbolUSDTBTC)
	{
		require.NoError(t, err)
	}
	{
		goldenFilename := "./testdata/future-common/v2-public-tickers.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
