package Prac2

import "fmt"

func ExamplePrac2() {
    list := [10]int{1, 4, 2, 30, 23, 13, 53, 3, 6, 16}
    list = sorting(list)
    fmt.Println(list)

    // Output:
    // [1 2 3 4 6 13 16 23 30 53]
}
