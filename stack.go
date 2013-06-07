package main

import ("fmt"
        "math/rand"
       )

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
    ret := (*s).St[len(s.St)-1]
    (*s).St = s.St[0:len(s.St)-1]
    return ret
}

func (s *Stack) peek() int {
    return (*s).St[len(s.St)-1]
}

func (s *Stack) length() int {
    return len((*s).St)
}

func main() {
    s := Stack{[]int{}}
    r := rand.Intn(15)
    fmt.Println("Stack init: ", s)
    fmt.Println("Stack length: ", s.length())
    for i:=0; i<r; i++ {
        s.push(i)
        fmt.Print("Stack push ", i)
        fmt.Println(": ", s.St)
    }
    fmt.Println("Stack length: ", s.length())
    for i:=len(s.St);i>0;i-- {
        fmt.Println("Stack peek: ", s.peek())
        fmt.Println("Stack pop: ", s.pop())
        fmt.Println("Stack: ", s.St)
    }
}
