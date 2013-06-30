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

func (d *DoubleLinkedList) add(n Node, str destination) (error) {
    err := fmt.Errorf("**ERROR** Invalid destination designated")
    return err
}

func (d *DoubleLinkedList) removeData(dat int) (error) {
    if !d.dataInList(dat) {
        err := fmt.Errorf("**ERROR** Data: %d not found in list",dat)
        return err
    }
}

func (d *DoubleLinkedList) dataInList(dat int) bool {
}

func (d *DoubleLinkedList) containsCycle() bool {
}

func (d *DoubleLinkedList) print() {
}

func main() {
    fmt.Println("Double linked list")
}