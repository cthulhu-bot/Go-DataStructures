package main

import (
        "testing"
       )

func Test_push_1 (t *testing.T) {
    s := Stack{[]int{}}
    s.push(1)
    s.push(2)
    ints := []int{1,2}
    int_test := true
    for i := range s.St {
        if s.St[i] != ints[i] {
            int_test = false
        }
    }
    if int_test {
        t.Log("push:    --test1 passed")
    } else {
        t.Error("Error in Stack.Push")
    }
}

func Test_pop_2 (t *testing.T) {
    s := Stack{[]int{1,2}}

    v,err := s.pop()
    if v == 2 && err == nil{
        t.Log("pop:     --test2 passed")
    } else if err != nil {
        t.Log(err)
    } else {
        t.Error("Error in Stack.Pop")
    }
}

func Test_peek_3 (t *testing.T) {
    s := Stack{[]int{5,6}}

    v,err := s.peek()
    if v == 6 && err == nil {
        t.Log("peek:    --test3 passed")
    } else if err != nil {
        t.Log(err)
    } else {
        t.Error("Error in Stack.Peek")
    }
}

func Test_length_4 (t *testing.T) {
    s := Stack{[]int{1,2,3,4,5}}
    if s.length() == 5 {
        t.Log("length:  --test4 passed")
    } else {
        t.Error("Error in Stack.Length")
    }
}