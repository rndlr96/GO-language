package main

import (
  "fmt"
)

type node struct {
  left *node
  val int
  right *node
}

func insert_tree(root *node, val int) *node {
  if root == nil {
    root = new(node)
    root.val = val
    root.left, root.right = nil, nil
  } else {
    if root.val < val {
      root.right = insert_tree(root.right, val)
    } else if root.val > val {
      root.left = insert_tree(root.left, val)
    }
  }
  return root
}

func print_tree(root *node) int {
    if root == nil {
      return 0
    } else {
      fmt.Println(root.val)
      print_tree(root.left)
      print_tree(root.right)
      return 0
    }
    return 0
}

func main(){
  var root *node

  root = insert_tree(root, 5)
  root = insert_tree(root, 3)
  root = insert_tree(root, 7)
  root = insert_tree(root, 1)
  root = insert_tree(root, 9)
  root = insert_tree(root, 6)

  print_tree(root)
}
