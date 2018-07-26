package main

import (
    "fmt"
    "sort"
    "strings"
)

func main() {
    list := []string{"Maria", "Andrew", "John"}
    sort.Sort(sort.StringSlice(list))

    var input_str string
    fmt.Printf("Please input searching string : ")
    fmt.Scanln(&input_str)

    for _, i := range list {
        if strings.Contains(i, input_str) == true {
            fmt.Printf("\"%s\" contains \"%s\"\n", i, input_str)
        }
    }
}
