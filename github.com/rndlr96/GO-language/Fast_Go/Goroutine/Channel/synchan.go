package package_name

import (
    "fmt"
//    "time"
)

func main() {
    done := make(chan bool)
    count := 3

    go func() {
        for i := 0 ; i < count ; i++ {
            fmt.Println("Goroutine : ", i)
            done <- true // waiting until main get done value
            //time.Sleep(1 * time.Second) // wait 1 Second
        }
    }()

    for i := 0 ; i < count ; i++ {
        <-done // waiting until value to come into the chan
        fmt.Println("Main : ", i)
    }
}
