package kata

import (
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	// your code
	arr := strconv.Itoa(n)
	decimal := make([]int, len(arr))

	for i, v := range arr {
		decimal[i] = int(v) - 48
	}

	sum := 0
	for i := 0; i < len(decimal); i++ {
		sum += int(math.Pow(float64(decimal[i]), float64(p+i)))
	}

	if sum < n || sum%n != 0 {
		return -1
	} else {
		return sum / n
	}
}
