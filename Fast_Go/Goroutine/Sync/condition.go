package package_name

import (
	"fmt"
	"runtime"
	"sync"
    "time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex)

	c := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func(n int) {
			mutex.Lock()
			c <- true
			fmt.Println("wait begin : ", n)
			cond.Wait()
			fmt.Println("wait end : ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c
	}

    // Signal
	for i := 0; i < 3; i++ {
		mutex.Lock()
		fmt.Println("signal : ", i)
		cond.Signal()
		mutex.Unlock()
	}

    time.Sleep(1 * time.Second)
    fmt.Printf("\n3초후 Broadcast를 통한 동기화가 시작됩니다.\n\n")
    time.Sleep(3 * time.Second)


    //Broadcast
    for i := 0; i < 3; i++ {
        go func(n int) {
            mutex.Lock()
            c <- true
            fmt.Println("wait begin : ", n)
            cond.Wait()
            fmt.Println("wait end : ", n)
            mutex.Unlock()
        }(i)
    }

    for i := 0; i < 3; i++ {
        <-c
    }

    mutex.Lock()
    fmt.Println("Broadcast")
    cond.Broadcast()
    mutex.Unlock()


	fmt.Scanln()

}
