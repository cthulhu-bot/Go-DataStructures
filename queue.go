package main

import "fmt"

type Queue struct {
    que []int
}

func (q *Queue) enqueue(i int) {
    (*q).que = append(q.que,i)
}

func (q *Queue) dequeue() (int, error) {
    if len(q.que) == 0 {
        return 0, errors.New("Attempted to dequeue an empty queue")
    }
    if len(q.que) == 1 {
        ret := (*q).que[0]
        (*q).que = {}
        return ret
    }

    ret := (*q).que[0]
    (*q).que = q.que[1:len(q.que)-1]
    return ret
}

func main() {
    q := Queue{[]int{}}

    fmt.Println("Enqueue: 2")
    q.enqueue(2)
    fmt.Println("Queue: ", q)
    for i:=len(q.que);i>0;i-- {
        fmt.Println("Dequeue: ", q.dequeue())
        fmt.Println("Queue: ", q)
    }
}