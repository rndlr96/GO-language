package unique

func FindUniq(arr []float32) float32 {
	// Do the magic
	key_map := make(map[float32]int)

	for _, i := range arr {
		key_map[i]++
	}

	for key, val := range key_map {
		if val == 1 {
			arr[0] = key
		}
	}
	return arr[0]
}
