package kata

func CountKprimes(k, start, end int)  []int {
    // your code
    var count, tmp, tmp2 int
    var prime []int
    for i := start ; i <= end ; i++ {
      tmp = i ; count = 0
      for ; tmp2 != tmp ; {
        tmp2 = tmp
        for j := 2 ; j < tmp ; j++ {
          if (tmp % j) == 0 && tmp != j {
            count++
            tmp = tmp / j
            break
          }
        }
      }
      if count == (k-1) { 
        prime = append(prime, i)
      }
    }
    return prime
}

func Puzzle(s int) int {
    // your code
    var count int
    onePrime := CountKprimes(1, 2, s)
    threePrime := CountKprimes(3, 2, s)
    sevenPrime := CountKprimes(7, 2, s)
    for _, i := range onePrime {
      for _, j := range threePrime {
        for _, k := range sevenPrime {
          if (i + j + k) == s { count++ }
        }
      }
    }
    return count
}