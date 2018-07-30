package triangle

import (
	"math"
)

func EquableTriangle(a, b, c int) bool {
	// Your code goes here
	sum := a + b + c
	s := float64(sum) / 2.0
	area := math.Sqrt(s * (s - float64(a)) * (s - float64(b)) * (s - float64(c)))
	if float64(sum) == area {
		return true
	} else {
		return false
	}
}
