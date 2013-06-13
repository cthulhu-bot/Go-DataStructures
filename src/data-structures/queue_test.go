package main

import (
        "testing"
)

func Test_enqueue_1 (t *testing.T) {
    q := Queue{[]int{}}
    q.enqueue(1)
    q.enqueue(2)

    testQueue := []int{1,2}
    testVal := true
    for i := range q.que {
        if q.que[i] != testQueue[i] {
            testVal = false
        }
    }
    if testVal {
        t.Log("enqueue:    --test1 passed")
    } else {
        t.Error("Error in queue.enqueue")
    }
}

func Test_dequeue_2 (t *testing.T) {
    q := Queue{[]int{1,2}}

    v,err := q.dequeue()
    if v == 1 && err == nil {
        t.Log("dequeue:    --test2 passed")
    } else if err != nil {
        t.Log(err)
    } else {
        t.Error("Error in queue.dequeue")
    }
}