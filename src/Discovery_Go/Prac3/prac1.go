// Package hangul provides basic functions for Hangul processing.
package hangul

import "fmt"

var (
	start = rune(44032) // "가"의 유니코드 포인트
	end   = rune(55204) // "힣" 다음 글자의 유니코드 포인트
)

// HasConsonantSuffix returns true if s has Hangul consonant jamo at the end.
func HasConsonantSuffix(s string) {
	numEnds := 28
	result := false
	for _, r := range s {
		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}
	if result == false {
        fmt.Printf("%s는 맛있다.\n", s)
    } else {
        fmt.Printf("%s은 맛있다.\n", s)
    }
}
