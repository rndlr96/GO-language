package kata

// ns : slice of indices
// xs, ys : chromosomes as slices of ints
func Crossover(ns []int, xs []int, ys []int) ([]int, []int) {
	// Your code here
	mymap := make(map[int]bool)

	for _, i := range ns {
		mymap[i] = true
	}

	for i, _ := range mymap {
		for j := i; j < len(xs); j++ {
			xs[j], ys[j] = ys[j], xs[j]
		}
	}
	return xs, ys
}
