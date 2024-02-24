package testhelper

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Compare :
func Compare(
	t *testing.T,
	want interface{},
	got interface{},
) {
	assert.JSONEq(t, string(convertToJSON(want)), string(convertToJSON(got)))
}

func convertToJSON(src interface{}) []byte {
	res, _ := json.MarshalIndent(src, "", "  ")
	return res
}

func Ptr[T any](v T) *T {
	return &v
}
