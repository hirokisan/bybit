//go:build integrationtestv5user

package integrationtestv5user

import (
	"testing"

	"github.com/dimkus/bybit/v2"
	"github.com/dimkus/bybit/v2/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestGetAPIKey(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	res, err := client.V5().User().GetAPIKey()
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-user-get-api-key.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
