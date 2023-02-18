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
