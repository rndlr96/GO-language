package main

import (
	"runtime"
)

// the return is `nil` if there is no buddy pair
func Buddy(start, limit int) []int {
	runtime.GOMAXPROCS(runtime.NumCPU())
	result := make(chan []int)
	var arr []int
	for j := 0; start+j < limit; j += 3000 {

		if start+j+3000 > limit {
			go func(j int) {
				for i := start + j; i <= limit; i++ {
					pair1 := sum(proper(i)) - 1
					pair2 := sum(proper(pair1)) - 1

					if pair2 == i {
						result <- []int{pair1, pair2}
					} else if i == limit {
						result <- nil
					}
				}
			}(j)
		} else {
			go func(j int) {
				for i := start + j; i < start+j+3000; i++ {
					pair1 := sum(proper(i)) - 1
					pair2 := sum(proper(pair1)) - 1

					if pair2 == i {
						result <- []int{pair1, pair2}
					}
				}
			}(j)
		}
	}
	arr = <-result

	return arr
	// your code
}

func proper(n int) []int {
	var prop []int

	for i := 1; 2*i <= n; i++ {
		if n%i == 0 {
			prop = append(prop, i)
		}
	}

	return prop
}

func sum(arr []int) int {
	var sum int

	for _, val := range arr {
		sum += val
	}

	return sum
}
