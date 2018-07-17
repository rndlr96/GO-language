package main

import
(
  "fmt"
)

var print = fmt.Println

func main ()
{
  var _map map[string]string = map[string]string
  {
    "name":     "Gu-Ik Jung",
    "country":  "Korea",
    "School" :  "Soongsil Univ"
  }

  print(<-map["name"])
}
