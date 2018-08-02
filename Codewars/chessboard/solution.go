package kata


func Game(n int) []int {
    // your code
    if n == 0 { return []int{0} }
  	sumSize := 2 * n -1
    sum := make([]int, sumSize+1)
    var result float64
    
  	for i := 1; i < n+1; i++ {
  		for j := 1; j < n+1; j++ {
  			sum[i+j-2] += j
  		}
    }
    
    for i, v := range sum {
      result += float64(v) / float64(i+2)
    }
    
    if int(n % 2) == 0 {
      return []int{int(result)}
    } else {
        sum[0] = int(result*2)
        sum[1] = 2
        return sum[:2]
    }
}