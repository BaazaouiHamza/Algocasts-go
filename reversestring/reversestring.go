package reversestring

import "strings"

func reverseString(str string) string {
	arr := strings.Split(str, "")
	rev := make([]string, 0)

	for i := len(arr) - 1; i >= 0; i-- {
		rev = append(rev, arr[i])
	}
	return strings.Join(rev, "")
}
