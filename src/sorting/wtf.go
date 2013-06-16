package main

import (
        "fmt"
)

func main() {
    foo := []int{5,12}
    bar := 3

    j := 0
    for j < len(foo) {
        if bar > foo[j] {
        } else if bar < foo[j] {
            frontHalf := append(foo[0:j],bar)
            backHalf := foo[j:len(foo)]

            fmt.Println("FrontHalf: ",frontHalf)
            fmt.Println("BackHalf: ",backHalf)
        }
        j++
    }

    fmt.Println(foo)
}