package Eval

import (
    "fmt"
)


func ExapmpleEval(){
    fmt.Println(Eval("5"))
    fmt.Println(Eval("1 + 2"))
    fmt.Println(Eval("1 - 2 + 3"))
    fmt.Println(Eval("3 * ( 3 + 1 * 3 ) / 2"))
    fmt.Println(Eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))
}
