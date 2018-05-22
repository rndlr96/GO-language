package main

import(
  "fmt"
)

func main() {

/*
=================for range================================
*/

  sum := 0
  for i := 1 ; i <= 10 ; i++{
    sum += i
    fmt.Println("operand+ :" ,sum)
  }

/*
=================for range================================
*/

  n := 1
  fmt.Println()
  for n < 10 {
    n *= 2
    fmt.Println("operand* :", n)
  }
/*
=================for range================================
*/

  names := []string{"Alex", "Ban", "Cool"}
  fmt.Println()
  for index, name := range names {
    fmt.Println("index : ",index," name : ", name)
  }

/*
===================break=================================
*/

  a := 1
  for a < 15 {
      if a == 5 {
        a += a
        continue
      }
      a++
      if a > 10 {
        break
      }
  }
  if a == 11{
    goto END
  }
  fmt.Println(a)

  END:
    fmt.Println("END")
}
