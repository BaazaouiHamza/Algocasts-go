package palindromes

import "strings"

func palindrome(str string) bool {
	sp := strings.Split(str, "")
	for i, j := 0, len(sp)-1; i < j; i, j = i+1, j-1 {
		sp[i], sp[j] = sp[j], sp[i]
	}
	return strings.Join(sp, "") == str
}
