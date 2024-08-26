package bybit

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateSyncTimeDelta(t *testing.T) {
	// given
	c := &Client{}
	remoteServerTimeRaw := "1688721231460000000"
	localTimestampNanoseconds := int64(1688721231560000000)
	expectedNsDelta := int64(100000000)
	recvWindowMs := int64(5000)
	nowTimestampMs := time.Now().UnixMilli()

	// when
	err := c.UpdateSyncTimeDelta(remoteServerTimeRaw, localTimestampNanoseconds)

	// then
	require.NoError(t, err)
	assert.Equal(t, expectedNsDelta, c.syncTimeDeltaNanoSeconds)

	timestampMsDelta := int64(math.Abs(float64(c.getTimestamp()) - float64(nowTimestampMs)))
	assert.Less(t, timestampMsDelta, recvWindowMs)
}
