package package_name

import "fmt"

func sum(a int, b int, c chan int){
    c <- a + b
}

func main() {
    c := make(chan int)

    go sum(1, 2, c)

    n := <- c
    fmt.Println(n)
}
