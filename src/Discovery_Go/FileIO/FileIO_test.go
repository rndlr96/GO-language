package TextIO

import (
	"fmt"
	"os"
	"strings"
)

func ExampleWriteTo() {
	lines := []string{
		"text1",
		"text2",
		"text3",
	}
	if err := WriteTo(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}
	// Output:
	// text1
	// text2
	// text3
}

func ExampleReadFrom() {
	r := strings.NewReader("text1\ntext2\ntext3\n")
	var lines []string
	if err := ReadFrom(r, &lines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [text1 text2 text3]
}
