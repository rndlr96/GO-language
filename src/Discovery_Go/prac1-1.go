package main

import (
	"fmt"
)

func price(i int) {
	for j := 1; j <= i; j++ {
		fmt.Println("apple : ", j*10, " orange : ", j*10+10)
	}
}

func main() {
	fmt.Print("Enter text: ")
	var input int
	fmt.Scanln(&input)
	price(input)
}
