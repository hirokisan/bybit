//go:build integrationtestv5position

package integrationtestv5position

import (
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestGetPositionInfo(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	symbol := bybit.SymbolV5BTCUSDT
	res, err := client.V5().Position().GetPositionInfo(bybit.V5GetPositionInfoParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   &symbol,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-position-get-position-info.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSetLeverage(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	res, err := client.V5().Position().SetLeverage(bybit.V5SetLeverageParam{
		Category:     bybit.CategoryV5Linear,
		Symbol:       bybit.SymbolV5BTCUSDT,
		BuyLeverage:  "1",
		SellLeverage: "1",
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-position-set-leverage.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSetTradingStop(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	category := bybit.CategoryV5Linear
	symbol := bybit.SymbolV5BTCUSDT
	qty := "0.01"
	{
		_, err := client.V5().Order().CreateOrder(bybit.V5CreateOrderParam{
			Category:  category,
			Symbol:    symbol,
			Side:      bybit.SideBuy,
			OrderType: bybit.OrderTypeMarket,
			Qty:       qty,
		})
		require.NoError(t, err)
	}
	price := "40000"
	res, err := client.V5().Position().SetTradingStop(bybit.V5SetTradingStopParam{
		Category:    category,
		Symbol:      symbol,
		PositionIdx: bybit.PositionIdxOneWay,
		TakeProfit:  &price,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-position-set-trading-stop.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
	{
		_, err := client.V5().Order().CreateOrder(bybit.V5CreateOrderParam{
			Category:  category,
			Symbol:    symbol,
			Side:      bybit.SideSell,
			OrderType: bybit.OrderTypeMarket,
			Qty:       qty,
		})
		require.NoError(t, err)
	}
}

func TestSwitchPositionMode(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()

	coin := bybit.CoinBTC
	res, err := client.V5().Position().SwitchPositionMode(bybit.V5SwitchPositionModeParam{
		Category: bybit.CategoryV5Inverse,
		Mode:     bybit.PositionModeBothSides,
		Coin:     &coin,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-position-switch-position-mode.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
