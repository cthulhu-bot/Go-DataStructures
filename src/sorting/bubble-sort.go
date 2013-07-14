package main

import (
        "fmt"
)

func bubbleSort(input []int) ([]int) {
    ptr := input[0]
    for i:=1; i < len(input); i++ {
        if i == len(input) {
            fmt.Println("Reached the end")
            return input
        }
        ptr++
    }
    return input
}

func main() {
    dat := []int{5,12,18,2,7,9,1,15,8,11}

    ret := bubbleSort(dat)
    fmt.Println("Final Sort: ", ret)
}