package main

import (
        "fmt"
)

type Node struct {
    data int
    next *Node
    prev *Node
}

type DoubleLinkedList struct {
    head *Node
    tail *Node
}

func main() {
    fmt.Println("Double linked list")
}