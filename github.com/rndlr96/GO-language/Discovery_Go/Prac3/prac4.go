package Prac4

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Create new Multiset and return this
func NewMultiSet() map[string]int {
	return make(map[string]int)
}

// Insert function is ADD val to Multiset m
func Insert(m map[string]int, val string) {
	var err error
    index := rand.Intn(99999)
	for m[strconv.Itoa(index)] != 0 {
		index = rand.Intn(99999)
	}

	if m[strconv.Itoa(index)], err = strconv.Atoi(val); err != nil {
		fmt.Println(err)
	}
}

// Erase function is DEL val from Multiset m
// if val not in Multiset, then nothing happens
func Erase(m map[string]int, val string) {
	// if no such element, delete is a no-op
    int_val,_ := strconv.Atoi(val)
	for key, m_val := range m {
        if m_val == int_val {
            delete(m, key)
            return
        }
    }
}

// Count function is Find the number of times val is included
// and return count
func Count(m map[string]int, val string) int {
	count := 0
	for _,num := range m {
        int_val,_ := strconv.Atoi(val)
		if num == int_val {
			count++
		}
	}
	return count
}

// String function is return string that Elements in {}
// separated by spaces
func String(m map[string]int) string {
	list := "{ "
	for _, val := range m {
		list += strconv.Itoa(val)
		list += " "
	}
	list += "}"
	return list
}
