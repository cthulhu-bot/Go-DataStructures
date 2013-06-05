package main

import "fmt"

type Stack struct {
    stack []int
    sp *int
}

func (s Stack) push(i int) {
    s.stack = append(s.stack, i)
}

func (s Stack) pop() int {
    return 5;
}

func (s Stack) peek() int {
    return 0;
}

func main() {
    s := Stack{}
    fmt.Println("Stack init: ", s)
    fmt.Println("Stack pop: ", s.pop())
    fmt.Println("Stack push: ", s)
}
