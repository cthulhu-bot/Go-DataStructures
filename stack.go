package main

import "fmt"

type Stack struct {
    stack []int
    sp *int
}

func (s Stack) push() {
}

func (s Stack) pop() int {
    return 0;
}

func (s Stack) peek() int {
    return 0;
}

func main() {
    s := Stack{}
    fmt.Println("Stack init: ", s)
}
