package main

import (
	"fmt"
)

type matrix struct {
	row    int
	column int
}

type optimal struct {
  start int
  k int
  end int
	val int
}

const (
  MaxInt = 999999
)
// parameter : matrix array for calculate multiple amount
// return : minimum multiple amount
func minMulti(price *[5][5]optimal, array []matrix, start int, end int, calc *int) int {

	count := 0
  min := MaxInt

  if start == end {
    return 0
  }

	for i := start; i < end; i++ {
		count = minMulti(price, array, start, i, calc) +
            minMulti(price, array, i+1, end, calc) +
            array[start].row * array[i].column * array[end].column

		*calc++
		if count < min {
			min = count
      price[start][end].start = start+1
      price[start][end].k = i+1
      price[start][end].end = end+1
			price[start][end].val = count
		}
	}

	return min
}

func main() {
	// generate matrix struct array and init each matrix row, column
	calc := 0

	var matrix_array = []matrix{
		matrix{
			row:    2,
			column: 5,
		},
		matrix{
			row:    5,
			column: 3,
		},
		matrix{
			row:    3,
			column: 6,
		},
		matrix{
			row:    6,
			column: 10,
		},
		matrix{
			row:    10,
			column: 7,
		},
	}

  price := [5][5]optimal{}

	fmt.Println("Minimum Calculate : ",minMulti(&price, matrix_array, 0, 4, &calc))
	fmt.Println("calculate count : ", calc)
	for i := 0 ; i < 5 ; i++ {
    for j := 0 ; j < 5 ; j++ {
      fmt.Print(price[i][j])
      fmt.Printf("\t")
    }
    fmt.Println()
  }

}
