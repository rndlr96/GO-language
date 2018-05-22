package implementation

import "fmt"

func zero(inval int)
{
  ival = 0
}

func zeroptr(iptr *int)
{
  *iptr = 0
}

func main()
{
  i := 1
  fmt.Println("intitial:", i)

  zero(i)
  fmt.Println("zero:", i)

  zeroptr(&i)
  fmt.Println("zeroptr:", i)

  fmt.Println("pointer:", &i)
}
