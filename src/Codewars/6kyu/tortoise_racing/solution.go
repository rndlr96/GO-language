package kata

func Race(v1, v2, g int) [3]int {
	// your code
	if v1 > v2 {
		return [3]int{-1, -1, -1}
	}
	var sec int

	speed := v2 - v1

	hour := int(g / speed)

	remain := (float64(g) / float64(speed)) - float64(hour)
	min := int(remain * 60)

	remain = (remain * 60) - float64(min)
	if (remain * 60) > float64(59) {
		min++
		sec = 0
	} else {
		sec = int(remain * 60)
		if (remain*60 - float64(sec)) > 0.9 {
			sec++
		}
	}

	return [3]int{hour, min, sec}
}
