package kata

import (
	"unicode"
	"fmt"
)

func Accum(s string) string {
	// your code
	var input string

	for k, v := range s {
		for i := 0; i < k+1; i++ {
			if i == 0 {
				input += string(unicode.ToUpper(v))
			} else {
				input += string(unicode.ToLower(v))
			}
		}
		input += string('-')
	}
	return input[:len(input)-1]
}
