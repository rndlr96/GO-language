package Interface

import "fmt"

func ExampleInterface() {
    fmt.Println(formatString(Person{"Maria", 20}))
    fmt.Println(formatString(&Person{"John", 12}))

    var andrew = new(Person)
    andrew.name = "Andrew"
    andrew.age = 35

    fmt.Println(formatString(andrew))

    // Output:
    // Maria 20
    // John 12
    // Andrew 35

}
