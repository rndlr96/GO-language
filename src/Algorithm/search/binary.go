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

func search_tree(root *node, val int) {
  if root == nil {
    fmt.Println("search fail")
  } else {
    if root.val == val {
      fmt.Println("search success")
    } else if root.val < val {
      root.right = search_tree(root.right, val)
    } else if root.val > val {
      root.left = search_tree(root.left, val)
    }
  }
}

func delete_tree(root *node, val int) {
  if root == nil {
    fmt.Println("Tree is empty")
  } else {
    if root.val == val {

    }
  }
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

  search_tree(root, 8)
  search_tree(root, 1)

}
