package kata

func ProductFib(prod uint64) [3]uint64 {
	// your code
	var result, find1, find2 uint64
	var tmp int

	for i := 0; result <= prod; i++ {
		find1 = findresult(i)
		find2 = findresult(i + 1)
		result = uint64(find1 * find2)
		if uint64(find1*find2) == prod {
			return [3]uint64{find1, find2, 1}
		}
		tmp = i
	}
	return [3]uint64{findresult(tmp), findresult(tmp + 1), 0}
}

func findresult(n int) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return findresult(n-1) + findresult(n-2)
	}
}
