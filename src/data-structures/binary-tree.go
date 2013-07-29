package main

import (
        "fmt"
)

type Node struct {
    data int
    leftChild *Node
    rightChild *Node
}

type BinaryTree struct {
    head *Node
}

func (b *BinaryTree) add(n Node) {
    fmt.Println("Eat a dick")
}

func (b *BinaryTree) isBalanced() (bool) {
    return false
}

func main() {
    tree := new (BinaryTree)
    n := Node{3,nil,nil}
    fmt.Println("add node: ", n)
    tree.add(n)
}