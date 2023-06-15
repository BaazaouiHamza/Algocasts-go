package palindromes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseFunctionExists(t *testing.T) {
	r := require.New(t)
	r.NotNil(palindrome, "Reverse function does not exist")
}

func TestPalindrome(t *testing.T) {
	r := require.New(t)
	r.False(palindrome(" aba"))
	r.False(palindrome("aba "))
	r.False(palindrome("greetings"))
	r.True(palindrome("1000000001"))
	r.False(palindrome("Fish hsif"))
	r.True(palindrome("pennep"))
}
