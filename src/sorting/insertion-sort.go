package main

import (
        "fmt"
)

func insertionSort(input []int) []int {
    fmt.Println("Input: ", input)
//    output := [len(input)]int{}
    for p:=1; p < len(input); p++ {
    }    
    
    return input
}

func main() {
    dat := []int{5,12,3,9,18}

    ret := insertionSort(dat)
    fmt.Println("Final Sort: ", ret)
}

