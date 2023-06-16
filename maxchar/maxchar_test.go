package maxchar

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxCharExists(t *testing.T) {
	r := require.New(t)
	r.NotNil(maxChar, "maxChar function does not Exist")
}

func TestMaxChar(t *testing.T) {
	r := require.New(t)
	r.Equal(string(maxChar("a")), "a")
	r.Equal("a", string(maxChar("abcdefghijklmnaaaaa")))
	r.Equal("1", string(maxChar("ab1c1d1e1f1g1")))
}
