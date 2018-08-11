package kata

import "math"

func ProductFib(prod uint64) [3]uint64 {
	// your code
	var result, find1, find2 uint64
	var tmp int

	for i := 2; result <= prod; i++ {
		find1 = ratio(i)
		find2 = ratio(i + 1)
		result = find1 * find2
		if result == prod {
			return [3]uint64{find1, find2, 1}
		}
		tmp = i
	}
	return [3]uint64{ratio(tmp), ratio(tmp + 1), 0}
}

func ratio(n int) uint64 {
	one := math.Pow(((1 + math.Sqrt(5.0)) / 2), float64(n))
	return uint64(math.Round(one / math.Sqrt(5.0)))
}
