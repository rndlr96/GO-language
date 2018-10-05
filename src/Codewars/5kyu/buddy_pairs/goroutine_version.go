package kata

import "math"

// the return is `nil` if there is no buddy pair
func Buddy(start, limit int) []int {
	result := make(chan []int)
	var arr []int

	for j := start; j <= limit; j++ {
		go func(j int) {
			pair1 := properSum(j) - 1
			pair2 := properSum(pair1) - 1

			if pair2 == j {
				result <- []int{pair2, pair1}
			}
		}(j)
	}

	arr = <-result
	return arr
	// your code
}

func properSum(n int) int {
	if n <= 0 {
		return 0
	}
	var sum int

	l := int(math.Sqrt(float64(n)))
	for i := 2; i <= l; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}

	return sum
}
