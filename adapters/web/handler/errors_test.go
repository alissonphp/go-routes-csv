package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Error Json"
	result := jsonError(msg)
	require.Equal(t, string([]byte(`{"message":"Error Json"}`)), string(result))
}
