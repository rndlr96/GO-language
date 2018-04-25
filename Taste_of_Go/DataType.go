package main

import "fmt"

func main() {
   var x float64
   x = 20.0
   fmt.Println(x)
   fmt.Printf("x is of type %T\n", x)

   var y float64 = 20.0

   z := 42 
   fmt.Println(y)
   fmt.Println(z)
   fmt.Printf("y is of type %T\n", y)
   fmt.Printf("z is of type %T\n", z)
}