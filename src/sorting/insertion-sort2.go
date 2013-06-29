package main

import (
        "fmt"
//        "errors"
)

func insertionSort(input []int) ([]int, error) {
    output := []int{}

    for i:=0; i < len(input); i++ {
        token := input[i]
        insertted := false
        j := 0
        for j < len(output) && !insertted {
            if token > output[j] {
                j++
            } else {
                fmt.Printf("i = %d j = %d\nToken = %d\n",i,j,token)
                frontHalf := output[0:j]
                backHalf := []int{}
                if len(output) > 0 {
                    backHalf = output[j+1:len(output)]
                }

                //Problematic
                frontHalf = append(frontHalf,token)
                insertted = true

                output = append(frontHalf,backHalf...)
            }
        }
    }

    return output,nil
}

func main() {
    input := []int{5,10,32,14,21}
    fmt.Println(insertionSort(input))
}