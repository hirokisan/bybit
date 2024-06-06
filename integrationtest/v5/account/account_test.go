//go:build integrationtestv5account

package integrationtestv5account

import (
	"testing"

	"github.com/dimkus/bybit/v2"
	"github.com/dimkus/bybit/v2/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestGetWalletBalance(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	res, err := client.V5().Account().GetWalletBalance(bybit.AccountTypeUnified, nil)
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-account-get-wallet-balance.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetAccountInfo(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	res, err := client.V5().Account().GetAccountInfo()
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-account-get-account-info.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetTransactionLog(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	limit := 1
	res, err := client.V5().Account().GetTransactionLog(bybit.V5GetTransactionLogParam{
		Limit: &limit,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-account-get-transaction-log.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetCollateralInfo(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	currency := "BTC"
	res, err := client.V5().Account().GetCollateralInfo(bybit.V5GetCollateralInfoParam{
		Currency: &currency,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-account-get-collateral-info.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSetCollateralCoin(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	coins := []bybit.Coin{bybit.CoinBTC}
	res, err := client.V5().Account().SetCollateralCoin(bybit.V5SetCollateralCoinParam{
		Coin:             bybit.CoinBTC,
		CollateralSwitch: bybit.CollateralSwitchV5On,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-account-set-collateral-coin.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
