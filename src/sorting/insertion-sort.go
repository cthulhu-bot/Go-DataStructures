package main

import (
        "fmt"
        "errors"
)

func insertionSort(input []int) ([]int, error) {
    if input == nil || len(input) == 1 {
        return nil, errors.New("**WARNING** input is either length 0 or 1")
    }

    fmt.Println("Input: ", input)
    output := []int{}
    output = append(output, input[0])
//    output := [len(input)]int{}
    for i:=1; i < len(input); i++ {
        tmp := input[i]
        insertted := false
        j := 0
        for j < len(output) && !insertted {
//            fmt.Println("Current output state: ",output)
//            fmt.Println("tmp to insert: ", tmp)
            if tmp > output[j] {
//                fmt.Printf("Compare input: %d\nSort position: %d\n",tmp,output[j])
                j++
            } else if tmp < output[j] {
                fmt.Println("Current output state: ",output)
                fmt.Println("Slice from j to length: ", output[j:len(output)])
//                fmt.Println("BackHalf: ", backHalf)
                fmt.Println("Current output state: ",output)
                backHalf := output[j:len(output)]
                fmt.Println("Current output state: ",output)
                fmt.Println("BackHalf: ", backHalf)
                fmt.Println("Current output state: ",output)

                // THIS LINE BREAKS SHIT AND CHANGES output (WTF)
                frontHalf := append(output[0:j],tmp)
//                fmt.Printf("Compare input: %d to sort position: %d\n",tmp,output[j])
                fmt.Println("Current output state: ",output)
                fmt.Println("FrontHalf: ", frontHalf)
                fmt.Println("Current output state: ",output)
                output = append(frontHalf,backHalf...)
                insertted = true
            }

            fmt.Println("temp output: ", output)
            if j+1 > len(output) {
//                fmt.Printf("Append %d to the end of the sorted list",tmp)
                output = append(output,tmp)
                break
            }

        }
        fmt.Println("Output: ",output)
    }    
    
    return output, nil
}

func main() {
    dat := []int{5,12,3}

    ret,err := insertionSort(dat)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Final Sort: ", ret)
    }
}

