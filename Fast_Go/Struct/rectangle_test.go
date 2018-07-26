package Rectangle

import (
    "fmt"
)

func ExampleRectangle() {
    rect1 := Rectangle{30, 30}
    rectangleScaleA(&rect1, 10)
    fmt.Println(rect1)

    rect2 := Rectangle{30, 30}
    rectangleScaleB(rect2, 10)
    fmt.Println(rect2)

    rect3 := Rectangle{30, 30}
    rect3.ScaleA(10)
    fmt.Println(rect3)

    rect4 := Rectangle{30, 30}
    rect4.ScaleB(10)
    fmt.Println(rect4)

    // Output:
    // {300 300}
    // {30 30}
    // {300 300}
    // {30 30}
}
