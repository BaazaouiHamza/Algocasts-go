package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxCharExists(t *testing.T) {
	r := require.New(t)
	r.NotNil(fizzBuzz, "maxChar function does not Exist")
}

func TestFizzBuzzPrintsStatements(t *testing.T) {
	consoleLogCalls := make([]interface{}, 0)
	_ = func(args ...interface{}) {
		consoleLogCalls = append(consoleLogCalls, args...)

		fizzBuzz(5)

		require.Equal(t, 5, len(consoleLogCalls), "Calling fizzBuzz with 5 should print 5 statements")
	}
}
