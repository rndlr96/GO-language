package kata

import "math"

func Gcdi(x, y int) int {

	if y == 0 {
		return int(math.Abs(float64(x)))
	} else {
		return Gcdi(y, x%y)
	}
}
func Som(x, y int) int {

	return x + y
}
func Maxi(x, y int) int {

	if x < y {
		return y
	} else {
		return x
	}
}
func Mini(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}

}
func Lcmu(x, y int) int {

	var lcm = (x * y) / Gcdi(x, y)
	if lcm < 0 {
		return lcm * -1
	} else {
		return lcm
	}
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) []int {

	for i := 0; i < len(arr); i++ {
		arr[i] = f(init, arr[i])
		init = arr[i]
	}
	return arr
}
