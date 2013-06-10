package main

import ("fmt"
        "math/rand"
        "errors"
       )

type Stack struct {
    St []int
}

func (s *Stack) push(i int) {
    (*s).St = append(s.St, i)
}

func (s *Stack) pop() (int,error) {
    if len(s.St) == 0 {
        return 0, errors.New("**ERROR** Attempted to pop an empty stack")
    }

    ret := (*s).St[len(s.St)-1]
    (*s).St = s.St[0:len(s.St)-1]
    return ret,nil
}

func (s *Stack) peek() (int,error) {
    if len(s.St) == 0 {
        return 0, errors.New("**ERROR** Stack is empty")
    }

    return (*s).St[len(s.St)-1],nil
}

func (s *Stack) length() int {
    return len((*s).St)
}

func foo() {
    s := Stack{[]int{}}
    r := rand.Intn(12)
    fmt.Println("Stack init: ", s)
    fmt.Println("Stack length: ", s.length())
    for i:=0; i<r; i++ {
        s.push(i)
        fmt.Print("Stack push ", i)
        fmt.Println(": ", s.St)
    }
    fmt.Println("Stack length: ", s.length())
    for i:=len(s.St);i>0;i-- {
        p,err := s.peek()
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println("Stack peek: ", p)
        }
        po,err := s.pop()
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println("Stack pop: ", po)
        }
        fmt.Println("Stack: ", s.St)
    }
    p,err := s.peek()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Stack peek: ", p)
    }
    po,err := s.pop()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Stack pop: ", po)
    }
    fmt.Println("Stack: ", s.St)
}
