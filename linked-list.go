package main

import (
        "fmt"
        "errors"
)

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}

func (l *LinkedList) add(n Node) {
    if (*l).head == nil {
        (*l).head = &n
        return
    }

    ptr := (*l).head
    for (*ptr).next != nil {
        ptr = (*ptr).next
    }
    (*ptr).next = &n
}

func (l *LinkedList) print() (error) {
    if (*l).head == nil {
        return errors.New("**WARNING** Printing out an empty list")
    }

    ptr := (*l).head
    for (*ptr).next != nil {
        fmt.Println("Node Data: ", (*ptr).data)
        ptr = (*ptr).next
    }
    fmt.Println("Node Data: ", (*ptr).data)
    return nil
}

func main() {
    list := new (LinkedList)
    fmt.Println("linked-list: ", list)
    n := Node{1,nil}
    fmt.Println("add node: ", n)
    list.add(n)
    n = Node{2,nil}
    fmt.Println("add node: ", n)
    list.add(n)

    err := list.print()
    if err != nil {
        fmt.Println(err)
    }
}
