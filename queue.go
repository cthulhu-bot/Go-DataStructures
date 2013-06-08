package main

import (
        "fmt"
        "errors"
)

type Queue struct {
    que []int
}

func (q *Queue) enqueue(i int) {
    (*q).que = append(q.que,i)
}

func (q *Queue) dequeue() (int, error) {
    if len(q.que) == 0 {
        return 0, errors.New("**ERROR** Attempted to dequeue an empty queue")
    }
    if len(q.que) == 1 {
        ret := (*q).que[0]
        (*q).que = []int{}
        return ret, nil
    }

    ret := (*q).que[0]
    (*q).que = q.que[1:len(q.que)-1]
    return ret, nil
}

func main() {
    q := Queue{[]int{}}

    fmt.Println("Enqueue: 2")
    q.enqueue(2)
    fmt.Println("Queue: ", q)
    for i:=len(q.que);i>0;i-- {
        d, err := q.dequeue()
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println("Dequeue: ", d)
            fmt.Println("Queue: ", q)
        }
    }
    d, err := q.dequeue()
    if err != nil {
       fmt.Println(err)
    } else {
        fmt.Println("Dequeue: ", d)
        fmt.Println("Queue: ", q)
    }
}