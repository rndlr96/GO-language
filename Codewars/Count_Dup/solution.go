package kata

import "strings"

func duplicate_count(s1 string) int {
    //your code goes here
    convS1 := strings.ToLower(s1)
    var exist = make(map[rune]int)
    var count int
    
    for _, i := range convS1 {
      exist[i]++
    }
    for _, v := range exist {
      if v > 1 { count++ }
    }
    return count
}

