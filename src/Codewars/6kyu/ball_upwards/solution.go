package kata

const G = 9.81

func MaxBall(v0 int) int {
	// your code
	var h, currentH float64

	for i := 0; i < v0; i++ {
		currentH = (float64(v0) * float64(i) / 36) - (0.5 * float64(i) * float64(i) * G / 100)
		if h <= currentH {
			h = currentH
		} else {
			return i - 1
		}
	}
	return 1
}
