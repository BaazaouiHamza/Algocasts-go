package reverseint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseIntFunctionExists(t *testing.T) {
	r := require.New(t)
	r.NotNil(reverseInt, "Reverse int function does not exist")
}

func TestReverseIntFlipsPositiveNumber(t *testing.T) {
	r := require.New(t)
	r.Equal(reverseInt(5), 5)
	r.Equal(reverseInt(15), 51)
	r.Equal(reverseInt(90), 9)
	r.Equal(reverseInt(2359), 9532)
}

func TestReverseIntFlipsNegativeNumber(t *testing.T) {
	r := require.New(t)
	r.Equal(reverseInt(-5), -5)
	r.Equal(reverseInt(-15), -51)
	r.Equal(reverseInt(-90), -9)
	r.Equal(reverseInt(-2359), -9532)
}
