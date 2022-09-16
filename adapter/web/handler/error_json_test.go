package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	mgs := "error message"
	e := jsonError(mgs)
	require.Equal(t, []byte(`{"message":"error message"}`), e)
}
