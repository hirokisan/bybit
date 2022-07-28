package testhelper

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Compare :
func Compare(
	t *testing.T,
	goldenFilename string,
	got []byte,
) {
	switch _, err := os.Stat(goldenFilename); err {
	case nil:
		want, err := os.ReadFile(goldenFilename)
		require.NoError(t, err)

		assert.JSONEq(t, string(want), string(got))
	case os.ErrNotExist:
		// noop
	default:
		t.Error(fmt.Errorf("unexpected error happened: %w", err))
	}
}

// SaveToFile :
func SaveToFile(name string, b []byte) error {
	return os.WriteFile(name, b, 0644)
}

// ConvertToJSON :
func ConvertToJSON(src interface{}) []byte {
	res, _ := json.MarshalIndent(src, "", "  ")
	return res
}

// UpdateFile :
func UpdateFile(t *testing.T, filename string, b []byte) {
	var updated bool
	updatedString, ok := os.LookupEnv("BYBIT_TEST_UPDATED")
	if ok && updatedString == "true" {
		updated = true
	}
	if updated {
		require.NoError(t, SaveToFile(filename, b))
		t.Log("golden file updated")
	}
}
