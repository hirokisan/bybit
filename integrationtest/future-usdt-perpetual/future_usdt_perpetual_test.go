//go:build integrationtestfutureusdtperpetual

package integrationtestfutureusdtperpetual

import (
	"testing"

	"github.com/hirokisan/bybit"
	"github.com/hirokisan/bybit/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestOrderBook(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual.OrderBook(bybit.SymbolInverseBTCUSD)
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-order-book-l2.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestTickers(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual.Tickers(bybit.SymbolInverseBTCUSD)
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

func TestSymbols(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual.Symbols()
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-symbols.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestOpenInterest(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual.OpenInterest(bybit.OpenInterestParam{
		Symbol: bybit.SymbolInverseBTCUSD,
		Period: bybit.Period1h,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-open-interest.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestBigDeal(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual.BigDeal(bybit.BigDealParam{
		Symbol: bybit.SymbolInverseBTCUSD,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-big-deal.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestAccountRatio(t *testing.T) {
	client := bybit.NewTestClient()
	limit := 10
	res, err := client.Future().USDTPerpetual.AccountRatio(bybit.AccountRatioParam{
		Symbol: bybit.SymbolInverseBTCUSD,
		Period: bybit.Period1h,
		Limit:  &limit,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-account-ratio.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
