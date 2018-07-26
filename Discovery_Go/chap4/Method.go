package Method

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strconv"
)

type MultiSet map[string]int
type SetOp func(m MultiSet, val string)

func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// Create new Multiset and return this
func NewMultiSet() MultiSet {
	return make(MultiSet)
}

// Insert function is ADD val to Multiset m
func (m MultiSet) Insert(val string){
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
func (m MultiSet) Erase(val string) {
	// if no such element, delete is a no-op
	int_val, _ := strconv.Atoi(val)
	for key, m_val := range m {
		if m_val == int_val {
			delete(m, key)
			return
		}
	}
}

// Count function is Find the number of times val is included
// and return count
func (m MultiSet) Count (val string) int {
	count := 0
	for _, num := range m {
		int_val, _ := strconv.Atoi(val)
		if num == int_val {
			count++
		}
	}
	return count
}

// String function is return string that Elements in {}
// separated by spaces
func (m MultiSet) String() string {
	list := "{ "
	for _, val := range m {
		list += strconv.Itoa(val)
		list += " "
	}
	list += "}"
	return list
}
