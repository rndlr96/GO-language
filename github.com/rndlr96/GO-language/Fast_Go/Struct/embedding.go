package Embedding

import (
    "fmt"
)

type Person struct {
    name string
    age int
}

type Student struct {
    p Person
    school string
    grade int
}

func (p *Person) greeting() {
    fmt.Println("Hello Person")
}

func (p *Student) greeting() {
    fmt.Println("Hello Student")
}
