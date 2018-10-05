package kata

import "math"

func ProductFib(prod uint64) [3]uint64 {
	// your code
	var result, find1, find2 uint64
	var tmp int

	for tmp = 2; result <= prod; tmp++ {
		find1 = ratio(tmp)
		find2 = ratio(tmp + 1)
		result = find1 * find2
		if result == prod {
			return [3]uint64{find1, find2, 1}
		}
	}
	return [3]uint64{ratio(tmp - 1), ratio(tmp), 0}
}

func ratio(n int) uint64 {
	one := math.Pow(((1 + math.Sqrt(5.0)) / 2), float64(n))
	return uint64(math.Round(one / math.Sqrt(5.0)))
}
