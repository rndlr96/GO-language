package Method

import (
	"fmt"
)

func ExampleMultiSet() {
	m := NewMultiSet()
	fmt.Println(m.String())
	fmt.Println(m.Count("3"))
	m.Insert("3")
	m.Insert("3")
	m.Insert("3")
	m.Insert("3")
	fmt.Println(m.String())
	fmt.Println(m.Count("3"))
	m.Insert("1")
	m.Insert("2")
	m.Insert("5")
	m.Insert("7")
	m.Erase("3")
	m.Erase("5")
	fmt.Println(m.Count("3"))
	fmt.Println(m.Count("1"))
	fmt.Println(m.Count("2"))
	fmt.Println(m.Count("5"))
	// Output:
	// { }
	// 0
	// { 3 3 3 3 }
	// 4
	// 3
	// 1
	// 1
	// 0
}
