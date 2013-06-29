package main

import (
        "fmt"
        "errors"
)

func insertionSort(input []int) ([]int, error) {
    if input == nil || len(input) == 1 {
        return nil, errors.New("**WARNING** input is either length 0 or 1")
    }

//    fmt.Println("Input: ", input)
    output := []int{}
    output = append(output, input[0])
    for i:=1; i < len(input); i++ {
        tmp := input[i]
        insertted := false
        j := 0
        for j < len(output) && !insertted {
            if tmp > output[j] {
                j++
            } else if tmp < output[j] {
                fmt.Println("Current output state: ",output)
                fmt.Println("Current tmp: ",tmp)
                backHalf := output[j:len(output)]
//                fmt.Println("Slice from j to len(output): ", backHalf)

                // THIS LINE BREAKS SHIT AND CHANGES output (WTF)
                // because the slice range from [x:x] = [] DUH
                frontHalf := []int{}
                if j == 0 {
                    //Insert token into front of slice
                    frontHalf = []int{tmp}
//                    fmt.Println("if BackHalf:  ", backHalf)
                } else {
                    //WHY THE FUCK WOULD THIS AFFECT BACKHALF
                  fmt.Println("BackHalf: ",backHalf)
//                    fmt.Println("j = ",j)
//                    fmt.Println("tmp = ",tmp)
                    frontHalf = append(output[0:j],tmp)
                    fmt.Printf("fronthalf = append(output[0:%d],%d)\n",j,tmp)
                fmt.Println("FrontHalf: ", frontHalf)
                   fmt.Println("BackHalf: ",backHalf)
                  fmt.Println("BackHalf: ",backHalf)
//                   fmt.Println("else BackHalf:  ", backHalf)
                }

//                fmt.Println("FrontHalf: ", frontHalf)
//                fmt.Println("BackHalf:  ", backHalf)
                output = append(frontHalf,backHalf...)
                insertted = true
            }

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
    dat := []int{7,5,4,9,8,1,2,3}

    ret,err := insertionSort(dat)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Final Sort: ", ret)
    }
}

