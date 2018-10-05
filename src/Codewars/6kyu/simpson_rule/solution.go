package kata

import "math"


func Simpson(n int) float64 {
    // your code
    var tmp float64
    b := math.Pi ; a := float64(0)
    
    h := (b - a) / float64(n)
    
    front := (b - a) / (3 * float64(n))
    
    first := f(a)
    
    second := f(b)
    
    for i := 1 ; i <= (n / 2) ; i++ {
      tmp += f(a + (float64(2 * i - 1)) * h)
    }
    third := 4 * tmp
    
    tmp = 0
    for i := 1 ; i <= ((n / 2) - 1) ; i++ {
      tmp += f(a + (2 * float64(i) * h))
    }
    forth := 2 * tmp
    
    return front * (first + second + third + forth)
    
    
    
}

func f(x float64) float64 {
  return 3 * math.Pow(math.Sin(x), 3) / 2
}
