//go:build integrationtestv5market

package integrationtestv5market

import (
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestGetKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.V5().Market().GetKline(bybit.V5GetKlineParam{
		Category: bybit.CategoryV5Spot,
		Symbol:   bybit.SymbolV5BTCUSDT,
		Interval: bybit.IntervalD,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-market-get-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetMarkPriceKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.V5().Market().GetMarkPriceKline(bybit.V5GetMarkPriceKlineParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   bybit.SymbolV5BTCUSDT,
		Interval: bybit.IntervalD,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-market-get-mark-price-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetIndexPriceKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.V5().Market().GetIndexPriceKline(bybit.V5GetIndexPriceKlineParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   bybit.SymbolV5BTCUSDT,
		Interval: bybit.IntervalD,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-market-get-index-price-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetPremiumIndexPriceKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.V5().Market().GetPremiumIndexPriceKline(bybit.V5GetPremiumIndexPriceKlineParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   bybit.SymbolV5BTCUSDT,
		Interval: bybit.IntervalD,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-market-get-premium-index-price-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetInstrumentsInfo(t *testing.T) {
	client := bybit.NewTestClient()
	{
		res, err := client.V5().Market().GetInstrumentsInfo(bybit.V5GetInstrumentsInfoParam{
			Category: bybit.CategoryV5Linear,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v5-market-get-instruments-info-inverse.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	}
	{
		res, err := client.V5().Market().GetInstrumentsInfo(bybit.V5GetInstrumentsInfoParam{
			Category: bybit.CategoryV5Option,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v5-market-get-instruments-info-option.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	}
	{
		res, err := client.V5().Market().GetInstrumentsInfo(bybit.V5GetInstrumentsInfoParam{
			Category: bybit.CategoryV5Spot,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v5-market-get-instruments-info-spot.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	}
}

func TestGetOrderbook(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.V5().Market().GetOrderbook(bybit.V5GetOrderbookParam{
		Category: bybit.CategoryV5Spot,
		Symbol:   bybit.SymbolV5BTCUSDT,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-market-get-orderbook.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetTickers(t *testing.T) {
	client := bybit.NewTestClient()
	{
		res, err := client.V5().Market().GetTickers(bybit.V5GetTickersParam{
			Category: bybit.CategoryV5Linear,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v5-market-get-tickers-inverse.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	}
	{
		coin := bybit.CoinBTC
		res, err := client.V5().Market().GetTickers(bybit.V5GetTickersParam{
			Category: bybit.CategoryV5Option,
			BaseCoin: &coin,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v5-market-get-tickers-option.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	}
	{
		res, err := client.V5().Market().GetTickers(bybit.V5GetTickersParam{
			Category: bybit.CategoryV5Spot,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v5-market-get-tickers-spot.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	}
}

func TestGetFundingRateHistory(t *testing.T) {
	client := bybit.NewTestClient()
	limit := 5
	res, err := client.V5().Market().GetFundingRateHistory(bybit.V5GetFundingRateHistoryParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   bybit.SymbolV5BTCUSDT,
		Limit:    &limit,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-market-get-funding-rate-history.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetPublicTradingHistory(t *testing.T) {
	client := bybit.NewTestClient()
	limit := 5
	res, err := client.V5().Market().GetPublicTradingHistory(bybit.V5GetPublicTradingHistoryParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   bybit.SymbolV5BTCUSDT,
		Limit:    &limit,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-market-get-public-trading-history.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetOpenInterest(t *testing.T) {
	client := bybit.NewTestClient()
	limit := 5
	res, err := client.V5().Market().GetOpenInterest(bybit.V5GetOpenInterestParam{
		Category:     bybit.CategoryV5Linear,
		Symbol:       bybit.SymbolV5BTCUSDT,
		IntervalTime: bybit.Period1h,
		Limit:        &limit,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-market-get-open-interest.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetHistoricalVolatility(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.V5().Market().GetHistoricalVolatility(bybit.V5GetHistoricalVolatilityParam{
		Category: bybit.CategoryV5Option,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-market-get-historical-volatility.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetInsurance(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.V5().Market().GetInsurance(bybit.V5GetInsuranceParam{})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-market-get-insurance.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
