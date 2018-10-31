package main

import (
  "fmt"
  "math"
  "os"
)

func main(){
  var input int

  fmt.Scanln(&input)

  if input < 1 || input > int(math.Pow(10,6)){
    os.Exit(1)
  }

  fmt.Println(count(input, 0))
}

func count(input, now int) int {
  var i int

  for i = 0 ; input != 1 ; i++ {
    if input % 3 == 0 && (input - 1) % 3 == 0 {
      return int(math.Min(float64(count(input, i)),float64(count(input-1, i))))
    } else if input % 3 == 0 {
      input = input / 3
    } else if (input-1) % 3 == 0 {
      input = (input-1) / 3
    } else if input % 2 == 0 {
      input = input / 2
    } else {
      input = input - 1
    }
  }
  return i+now
}
