package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func random(n int) {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	fmt.Println(n)
}

func main() {

	// Single Core
	for i := 0; i < 10; i++ {
		go random(i)
	}

	//Multi Core
	runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.GOMAXPROCS  : Set the number of CPUs to use
	// runtime.NumCPU      : Get the number of cpu

	fmt.Println("Core count : ", runtime.GOMAXPROCS(0))
	// Print setting value

	s := "Hello World"

	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}

    // 고루틴을 종료하려면 함수 안에서 return으로 빠져나오거나
    // runtime.Goexit 함수를 사용해야 한다. 단 리턴값은 사용할 수 없다.

    // if you want exit from Goroutine, then you must
    // use return in function or use runtime.Goexit()
    // you can't use return value
