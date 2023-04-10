//go:build integrationtestv5account

package integrationtestv5account

import (
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest/testhelper"
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
	res, err := client.V5().Account().GetAccountInfo(1, MarginModeRegular, "")
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-account-get-account-info.json" // TODO
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
