//go:build integrationtestv5user

package integrationtestv5user

import (
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestGetPositionInfo(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	res, err := client.V5().User().GetAPIKeyInfo()
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-user-get-api-key-info.json" // TODO
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
