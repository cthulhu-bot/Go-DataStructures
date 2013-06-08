package stack

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
        t.Log("push:  --test1 passed")
    } else {
        t.Error("Error in Stack.Push")
    }
}
