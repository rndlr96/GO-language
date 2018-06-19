package main

import (
  "fmt"
)

func main(){
  var name string
  category := 1

  switch category {
  case 1:
    name = "Paper Book"
  case 2:
    name = "eBook"
  case 3,4:
    name = "Blog"
  default:
    name = "Other"
  }

  fmt.Println(name)
}
