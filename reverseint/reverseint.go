package reverseint

import (
	"strconv"
	"strings"
)

func reverseInt(n int) int {
	rev := strings.Split(strconv.Itoa(n), "")
	if rev[0] == "-" {
		rev = append(rev[:0], rev[1:]...)
	}
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	revInt, err := strconv.ParseInt(strings.Join(rev, ""), 10, 32)
	if err != nil {
		panic(err)
	}
	sign := 1
	if n < 0 {
		sign = -1
	}
	return int(revInt) * sign

}
