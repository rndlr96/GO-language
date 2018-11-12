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

func init_matrix(price *[5][5]optimal){
  for i := 0 ; i < 5 ; i++ {
    price[i][i].val = 0;
  }
}

func main() {

	calc, count := 0, 0

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

  for i := 1 ; i < 5 ; i++ {
    for j := 0 ; j < 5 - i ; j++ {
      limit := i + j
      min := MaxInt

      for k := j ; k < limit ; k++ {
        count = price[j][k].val + price[k+1][limit].val +
                matrix_array[j].row * matrix_array[k].column * matrix_array[limit].column
        calc++
        if count < min {
          min = count
          price[j][limit].val = min
        }
      }
    }
  }

	fmt.Println("calculate count : ", calc)
  for i := 0 ; i < 5 ; i++ {
    for j := 0 ; j < 5 ; j++ {
      fmt.Print(price[i][j])
      fmt.Printf("\t")
    }
    fmt.Println()
  }
}
