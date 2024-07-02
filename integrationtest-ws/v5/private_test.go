//go:build integrationtestwsv5

package integrationtestwsv5

import (
	"context"
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest-ws/testhelper"
	"github.com/stretchr/testify/require"
)

func TestV5Private_Position(t *testing.T) {
	wsClient := bybit.NewTestWebsocketClient().WithAuthFromEnv()
	svc, err := wsClient.V5().Private()
	require.NoError(t, err)

	require.NoError(t, svc.Subscribe())

	_, err = svc.SubscribePosition(
		func(response bybit.V5WebsocketPrivatePositionResponse) error {
			goldenFilename := "./testdata/private-v5-position.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(response))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(response))
			return nil
		},
	)
	require.NoError(t, err)

	require.NoError(t, svc.Start(context.Background(), nil))
	// When you make a change to a position, the change is recorded.
	// After recorded, please stop manually.
}

func TestV5Private_Execution(t *testing.T) {
	wsClient := bybit.NewTestWebsocketClient().WithAuthFromEnv()
	svc, err := wsClient.V5().Private()
	require.NoError(t, err)

	require.NoError(t, svc.Subscribe())

	_, err = svc.SubscribeExecution(
		func(response bybit.V5WebsocketPrivateExecutionResponse) error {
			goldenFilename := "./testdata/private-v5-execution.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(response))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(response))
			return nil
		},
	)
	require.NoError(t, err)

	require.NoError(t, svc.Start(context.Background(), nil))
}

func TestV5Private_Wallet(t *testing.T) {
	wsClient := bybit.NewTestWebsocketClient().WithAuthFromEnv()
	svc, err := wsClient.V5().Private()
	require.NoError(t, err)

	require.NoError(t, svc.Subscribe())

	_, err = svc.SubscribeWallet(
		func(response bybit.V5WebsocketPrivateWalletResponse) error {
			goldenFilename := "./testdata/private-v5-wallet.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(response))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(response))
			return nil
		},
	)
	require.NoError(t, err)

	require.NoError(t, svc.Start(context.Background(), nil))
	// When you make a change to a position, the change is recorded.
	// After recorded, please stop manually.
}
