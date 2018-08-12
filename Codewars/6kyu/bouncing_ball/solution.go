package kata

func BouncingBall(h, bounce, window float64) int {
	// your code
	if h <= 0.0 || bounce <= 0.0 || bounce >= 1.0 || window >= h {
		return -1
	}
	count := 1
	h *= bounce

	for h > window {
		count += 2
		h *= bounce
	}
	return count
}
