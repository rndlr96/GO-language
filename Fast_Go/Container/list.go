package main

import (
	"container/list"
	"fmt"
)

func main() {

	list1 := list.New()
	list1.PushBack(10)
	list1.PushBack(20)
	list1.PushBack(30)

    list2 := list.New()
    list2.PushBack(40)
    list2.PushBack(50)

    list1.PushBackList(list2)

	fmt.Println("Front ", list1.Front().Value, "\n")
	fmt.Println("Back ", list1.Back().Value, "\n")


    fmt.Println("Front -> Back")
	for e := list1.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
    fmt.Println()


    fmt.Println("Back -> Front")
    for e := list1.Back() ; e != nil ; e = e.Prev() {
        fmt.Println(e.Value)
    }
    fmt.Println()


    fmt.Println("list len : ", list1.Len())
}
