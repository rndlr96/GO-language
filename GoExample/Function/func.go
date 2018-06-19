package main

import(
  "fmt"
)

func main() {
  print(0)
  msg := "Hello"
  val_say(msg)

  print(1)
  ref_say(&msg)
  fmt.Println("after msg : ",msg)

  print(2)
  var_say("one", "two", "three", "four")
  fmt.Println()
  var_say("one")

  print(3)
  sum := return_say(1,3,5,7,9)
  fmt.Println("sun : ", sum)

  print(4)
  count, sum := mul_return_say(1,3,5,7,9)
  fmt.Println("count : ", count," sum : ", sum)
}


func print(num int){
  fmt.Println()

  switch num{
  case 0:
    fmt.Println("========== Call By Value =================")
  case 1:
    fmt.Println("========== Call By Reference =============")
  case 2:
    fmt.Println("========== Variadic Function =============")
  case 3:
    fmt.Println("========== Return Function ===============")
  case 4:
    fmt.Println("========== Multi Return Fuction ==========")
  }

  fmt.Println()
}

func val_say(msg string){
  fmt.Println("var_say func: ", msg)
}

func ref_say(msg *string){
  fmt.Println("before msg : ",*msg)
  *msg = "Changed"
}

func var_say(msg ...string){
  for index, s := range msg {
    fmt.Println("index : ", index, " string : ", s)
  }
}

func return_say(nums ...int) int {
  s := 0
  for _, n := range nums {
    s += n
  }
  return s
}

func mul_return_say(nums ...int) (int, int) {
  s := 0
  count := len(nums)
  for _, n := range nums {
    s += n
  }
  return count, s

}
