package main

import (
  "fmt"
)

func main() {
  var var1 string = "Jung"
  fmt.Println(var1);
  var var2 = "Gu" // Type can be omitted.
  fmt.Println(var2);
  var3 := "Ik"    // when we use ':' var can be omitted
  fmt.Println(var3);

  var ( // Can declare at once
    site    string = "github.com"
    name    string = "Pandog"
    job     string = "blockchain Core Developer"
  )

    fmt.Println(site);
    fmt.Println(name);
    fmt.Println(job);
}
