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

func (d *DoubleLinkedList) add(n Node) {
}

func (d *DoubleLinkedList) removeData(dat int) (error) {
}

func (d *DoubleLinkedList) dataInList(dat int) bool {
}

func (d *DoubleLinkedList) containsCycle() bool {
}

func main() {
    fmt.Println("Double linked list")
}