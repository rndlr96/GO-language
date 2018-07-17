package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	/*
		say("Sync")

		go say("Async1")
		go say("Async2")
		go say("Async3")

		time.Sleep(time.Second * 3)
	*/

	/* anonymous function
	   var wait sync.WaitGroup //
	   wait.Add(2)

	   go func(){
	     defer wait.Done()
	     fmt.Println("Hello")
	   }

	   go func(msg string){
	     defer wait.Done()
	     fmt.Println(msg)
	   }("Hi")

	   wait.Wait()
	*/
	runtime.GOMAXPROCS(4) // for multi CPU
}
