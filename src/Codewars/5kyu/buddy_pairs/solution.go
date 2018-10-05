package kata

import "math"

// the return is `nil` if there is no buddy pair
func Buddy(start, limit int) []int {

	for i := start; i <= limit; i++ {
		pair1 := properSum(i) - 1
		pair2 := properSum(pair1) - 1

		if i < pair1 && pair2 == i {
			return []int{pair2, pair1}
		}
	}

	return nil
	// your code
}

func properSum(n int) int {
	if n <= 0 {
		return 0
	}
	sum := 1

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
