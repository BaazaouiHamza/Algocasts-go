package maxchar

import (
	"fmt"
)

func maxChar(str string) rune {

	char := make(map[rune]int)

	for _, c := range str {
		char[c] = char[c] + 1
	}

	keys := make([]rune, 0, len(char))
	for k, v := range char {
		keys = append(keys, k)
		fmt.Printf("Key:%s,Value:%v \n", string(k), v)
	}

	maxKey := keys[0]
	for i := 1; i < len(keys); i++ {

		if char[keys[i]] > char[maxKey] {
			maxKey = keys[i]
		}
	}

	return maxKey

}
