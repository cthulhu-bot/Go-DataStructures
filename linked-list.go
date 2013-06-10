package main

import (
        "fmt"
)

type Node struct {
    data int
    next *Node
}

func main() {
    n := Node{0,nil}
    fmt.Println("linked-list")
}
