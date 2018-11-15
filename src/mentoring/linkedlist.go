package main

import "fmt"

type node struct {
  val int
  point *node
}

func insert(root *node, val int) *node {
  if root == nil {
    root = new(node)
    root.val = val
    root.point = nil
  } else {
    root.point = insert(root.point, val)
  }
  return root
}

func search(root *node, val int) {
  count := 1
  if root == nil {
    fmt.Println("list is empty")
  } else {
    if root.val == val {
      fmt.Println("search value :", val, " result : success")
    } else {
      count++
      if root.point == nil {
        fmt.Println("search value :", val, " result : fail")
      } else {
        search(root.point, val)
      }
    }
  }
}

func delete(root *node, val int) *node {
  if root == nil {
    fmt.Println("list is empty")
  } else {
    if root.val == val {
      return root.point
    } else {
      if root.point == nil {
        fmt.Println("value is not in list")
      } else {
        if root.point.val == val {
          root.point = root.point.point
        } else {
          root.point = delete(root.point, val)
        }
      }
      return root
    }
  }
  return root
}

func print(root *node) {
  if root == nil {
    fmt.Println("list is empty")
  } else {
    for ;root.point != nil; root = root.point {
      fmt.Printf("%d -> ", root.val)
    }
    fmt.Printf("%d", root.val)
    fmt.Println()
  }
}

func main(){
  var root *node;

  root = insert(root, 3)
  root = insert(root, 5)
  root = insert(root, 7)
  print(root)

  search(root, 5)

  root = delete(root, 5)
  print(root)

  search(root, 5)
}
