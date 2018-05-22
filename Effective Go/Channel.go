// Channel 고루틴이 서로 통신하고 변수를 동기화할 수 있는 수단
// Go의 concurrency 관련해서 학습 필요

package main

import (
  "fmt"
  "os"
)

var print = fmt.Println

func add(a int, b int, c chan int)
{
  c <- a + b
}

func main()
{
  var c = make(chan int)
  go add(1, 3, c) // start goroutin
  print(<-c)

  var b []byte = make([]byte, 1)
  os.Stdin.Read(b)
}
