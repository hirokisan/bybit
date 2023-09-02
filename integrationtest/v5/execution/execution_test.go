//go:build integrationtestv5execution

package integrationtestv5execution

import (
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestGetexecutionList(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	symbol := bybit.SymbolV5BTCUSDT
	limit := 1
	res, err := client.V5().Execution().GetExecutionList(bybit.V5GetExecutionParam{
		Category: bybit.CategoryV5Linear,
		Symbol:   &symbol,
		Limit:    &limit,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-execution-get-execution-list.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
