package kata

import (
	"sort"
	"strconv"
)

func Decomp(n int) string {
	// your code
	mymap := make(map[int]int)
	var result string

	for i := 2; i <= n; i++ {
		tmp := i
		for j := 2; j <= i; j++ {
			if tmp%j == 0 {
				mymap[j] += 1
				tmp = tmp / j
				j--
			}
		}
	}

	keys := make([]int, 0, len(mymap))
	for k := range mymap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, key := range keys {
		if mymap[key] == 1 {
			result += strconv.Itoa(key) + " * "
		} else {
			result += strconv.Itoa(key) + "^" + strconv.Itoa(mymap[key]) + " * "
		}
	}

	return result[:len(result)-3]
}
