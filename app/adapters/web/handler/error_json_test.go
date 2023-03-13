package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// suíte de testes unitários da função jsonError

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello Json"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"message":"Hello Json"}`), result)
}
