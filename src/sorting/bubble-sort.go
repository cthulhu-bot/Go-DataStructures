package main

import (
        "fmt"
)

func bubbleSort(input []int) ([]int) {
    ptr := &input[0]
    // Wont work with negative inputs
    lastNum := 0
    for i:=0; i <= len(input); i++ {
        if i == len(input) {
            fmt.Println("Reached the end")
            return input
        } else if input[i] > lastNum {
            *ptr = input[i]
            input[i] = lastNum
        }
        lastNum = input[i]
        ptr = &input[i]
    }
    return input
}

func isSorted() (bool) {
    return false
}

func main() {
    dat := []int{5,12,18,2,7,9,1,15,8,11}

    ret := bubbleSort(dat)
    fmt.Println("Final Sort: ", ret)
}