package main

import (
        "fmt"
)

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}

func (l *LinkedList) add(n Node) {
}

func main() {
    list := new (LinkedList)
//    n := Node{0,}

//    list.add(n)
    fmt.Println("linked-list: ", list)
}
