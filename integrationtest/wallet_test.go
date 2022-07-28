//go:build integrationtest

package integrationtest

import (
	"testing"

	"github.com/hirokisan/bybit"
	"github.com/hirokisan/bybit/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestBalance(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Wallet().Balance(bybit.CoinUSDT)
		{
			require.NoError(t, err)
			require.Equal(t, "OK", res.RetMsg)
		}
		{
			goldenFilename := "./testdata/v2-private-wallet-balance.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Wallet().Balance(bybit.CoinBTC)
		require.Error(t, err)
	})
}
