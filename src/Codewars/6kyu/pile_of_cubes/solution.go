package kata

import "math"

func FindNb(m int) int {
  // your code
  var tmp float64
  sqrt := math.Sqrt(float64(m))
  if sqrt != math.Floor(sqrt) {
    return -1
  }
  start := math.Floor(math.Sqrt(sqrt))
  sqrt = sqrt * 2
  for i := int(start) ; tmp < sqrt ; i++ {
    tmp = float64(i) * (float64(i)+1)
    if tmp == sqrt {
      return i
    }
  }
  return -1
}