package reversestring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseFunctionExists(t *testing.T) {
	r := require.New(t)
	r.NotNil(reverseString, "Reverse function does not exist")
}

func TestReverseReversesString(t *testing.T) {
	r := require.New(t)
	result := reverseString("abcd")
	expected := "dcba"
	r.Equal(expected, result, "Reverse result is incorrect")
}

func TestReverseReversesStringWithLeadingSpaces(t *testing.T) {
	r := require.New(t)
	result := reverseString("  abcd")
	expected := "dcba  "
	r.Equal(expected, result, "Reverse result is incorrect")
}
