package main

import (
	"fmt"
    "os"
)

func main() {
    // os.Open has two reutrn values.
    // one is file object and
    // the other value is error { err value is nil = complete }
	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
	}
    // The defer registers a function to be called
    // if it is outside the function.
    defer f.Close()
    var num int
	if _, err := fmt.Fscanf(f, "%d\n", &num); err == nil {
        fmt.Println(num)
	}

    f, err = os.Create("data.txt")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("Insert some number : ")
    fmt.Scanf("%d\n", &num)
    if _, err := fmt.Fprintf(f, "%d\n", num); err == nil {
        fmt.Println(num)
        return
    }
}
