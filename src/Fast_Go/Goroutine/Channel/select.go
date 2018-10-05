package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan int)
    c2 := make(chan string)

    go func() {
        for i := 0 ; i < 20 ; i++ {
            c1 <- 10
            time.Sleep(100 * time.Millisecond)
        }
    }()

    go func() {
        for i := 0 ; i < 4 ; i++ {
            c2 <- "Hello world"
            time.Sleep(500 * time.Millisecond)
        }
    }()

    go func() {
        for {
            select {
            case i := <- c1 :
                fmt.Println("c1 : ",i)
            case j := <- c2 :
                fmt.Println("c2 : ",j)
            }
        }
    }()

    time.Sleep(10 * time.Second)
}
