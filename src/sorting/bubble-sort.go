package main

import (
        "fmt"
)

type BubbleSort struct {
}

func (b *BubbleSort) sort(input []int) ([]int) {
    for !isSorted(input) {
        ptr := &input[0]
        // Wont work with negative inputs
        lastNum := 0
        for i:=0; i < len(input); i++ {
            if i == len(input) {
                fmt.Println("Reached the end")
                return input
            } else if input[i] < lastNum {
                *ptr = input[i]
                input[i] = lastNum
            }
            lastNum = input[i]
            ptr = &input[i]
        }
    }
    return input
}

func isSorted(input []int) (bool) {
    ptr := &input[0]
    for i:=1; i < len(input); i++ {
        curr := input[i]
        if curr < *ptr {
            return false
        }
        ptr = &input[i]
    }
    return true
}

func main() {
    dat := []int{5,12,18,2,7,9,1,15,8,11}

    b := BubbleSort{}

    ret := b.sort(dat)
    fmt.Println("Final Sort: ", ret)
}