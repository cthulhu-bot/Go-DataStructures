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

func (l *LinkedList) remove(dat int) {

    if (*l).head.data == dat {
        (*l).head = (*l).head.next
    }

    ptr := (*l).head
    for (*ptr).next != nil {
        fmt.Println("ptr.next: ",(*ptr).next)
        fmt.Println("ptr.next.data: ",(*ptr).next.data)
        if (*ptr).next.data == dat && (*ptr).next.next == nil {
            (*ptr).next = nil
        }
        if (*ptr).next.data == dat && (*ptr).next.next != nil {
            (*ptr).next = (*ptr).next.next
        }
        ptr = (*ptr).next
    }
}

func (l *LinkedList) print() (error) {
    if (*l).head == nil {
        return errors.New("**WARNING** Printing out an empty list")
    }

    ptr := (*l).head
    fmt.Println("Linked List: ")
    for (*ptr).next != nil {
        fmt.Printf("Node %d -> ", (*ptr).data)
        ptr = (*ptr).next
    }
    fmt.Println("Node", (*ptr).data)
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
    n = Node{3,nil}
    fmt.Println("add node: ", n)
    list.add(n)
    n = Node{4,nil}
    fmt.Println("add node: ", n)
    list.add(n)

    err := list.print()
    if err != nil {
        fmt.Println(err)
    }

    list.remove(4)
    err = list.print()
    if err != nil {
        fmt.Println(err)
    }
}
