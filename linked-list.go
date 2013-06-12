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

func (l *LinkedList) dataInList(dat int) bool {
    ptr := (*l).head

    for (*ptr).next != nil {
        if (*ptr).data == dat { return true }
        ptr = (*ptr).next
    }
    if (*ptr).data == dat { return true }
    return false
}

func (l *LinkedList) removeData(dat int) (error) {
    if !l.dataInList(dat) {
        err := fmt.Errorf("**ERROR** Data: %d not found in list",dat)
        return err
    }

    if (*l).head.data == dat {
        (*l).head = (*l).head.next
    }

    ptr := (*l).head
    for (*ptr).next != nil {
        if (*ptr).next.data == dat && (*ptr).next.next != nil {
            (*ptr).next = (*ptr).next.next
        } else if (*ptr).next.data == dat && (*ptr).next.next == nil {
            (*ptr).next = nil
            return nil
        }
        ptr = (*ptr).next
    }
    return nil
}

func (l *LinkedList) print() {
    if (*l).head == nil {
        fmt.Println("**WARNING** Printing out an empty list")
        return
    }

    ptr := (*l).head
    fmt.Println("Linked List: ")
    for (*ptr).next != nil {
        fmt.Printf("Node %d -> ", (*ptr).data)
        ptr = (*ptr).next
    }
    fmt.Println("Node", (*ptr).data)
    return
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

    list.print()

    err := list.removeData(3)
    list.print()
    err = list.removeData(5)
    if err != nil {
        fmt.Println(err)
    }
}
