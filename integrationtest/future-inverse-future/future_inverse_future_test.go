//go:build integrationtestfutureinversefuture

package integrationtestfutureinversefuture

import (
	"testing"
	"time"

	"github.com/hirokisan/bybit"
	"github.com/hirokisan/bybit/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestOrderBook(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().OrderBook(bybit.SymbolInverseBTCUSD)
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

func TestListKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().ListKline(bybit.ListKlineParam{
		Symbol:   bybit.SymbolInverseBTCUSD,
		Interval: bybit.Interval120,
		From:     int(time.Now().AddDate(0, 0, -1).Unix()),
	})
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-kline-list.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestTickers(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().Tickers(bybit.SymbolInverseBTCUSD)
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

func TestTradingRecords(t *testing.T) {
	client := bybit.NewTestClient()
	limit := 10
	res, err := client.Future().InverseFuture().TradingRecords(bybit.TradingRecordsParam{
		Symbol: bybit.SymbolInverseBTCUSD,
		Limit:  &limit,
	})
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-trading-records.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSymbols(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().Symbols()
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

func TestIndexPriceKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().IndexPriceKline(bybit.IndexPriceKlineParam{
		Symbol:   bybit.SymbolInverseBTCUSD,
		Interval: bybit.IntervalD,
		From:     int(time.Now().AddDate(0, 0, -1).Unix()),
	})
	{
		require.NoError(t, err)
		require.Equal(t, "OK", res.RetMsg)
	}
	{
		goldenFilename := "./testdata/v2-public-index-price-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestOpenInterest(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().OpenInterest(bybit.OpenInterestParam{
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
	res, err := client.Future().InverseFuture().BigDeal(bybit.BigDealParam{
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
	res, err := client.Future().InverseFuture().AccountRatio(bybit.AccountRatioParam{
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
