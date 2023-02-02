package testingcommon

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/stretchr/testify/require"
)

// Converts map created by json decoder to struct
// out should be pointer (&payload)
func ExtractApiSuccessResponse[T any](t *testing.T, bytesBuffer *bytes.Buffer) *httpo.ApiSuccessResponse[T] {
	var response *httpo.ApiSuccessResponse[T]
	err := json.NewDecoder(bytesBuffer).Decode(&response)
	require.Nil(t, err)
	return response
}
