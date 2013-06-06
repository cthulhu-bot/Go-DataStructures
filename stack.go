package main

import "fmt"

type Stack struct {
    St []int
}

func newStack(foo []int, bar *int) *Stack {
    return &Stack{foo}
}

func (s *Stack) push(i int) {
    (*s).St = append(s.St, i)
}

func (s *Stack) pop() int {
    return 0;
}

func (s *Stack) peek() int {
    return (*s).St[len(s.St)]
}

func (s *Stack) length() int {
    return len((*s).St)
}

func main() {
    s := Stack{[]int{1,2}}
    fmt.Println("Stack init: ", s)
    fmt.Println("Stack length: ", s.length())
    i := 3
    s.push(i)
    fmt.Print("Stack push ", i)
    fmt.Println(": ", s.St)
}
