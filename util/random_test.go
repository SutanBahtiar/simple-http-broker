package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	r := RandomString(2)
	require.NotZero(t, len(r))
}
