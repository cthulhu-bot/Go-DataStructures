package main

import (
        "fmt"
)

func main() {
    foo := []int{5,12}
    bar := 3
    ret := []int{}

    j := 0
    fmt.Println("Initial foo: ",foo)
    for j < len(foo) {
        if bar < foo[j] {
            fmt.Println("foo: ",foo)
            fmt.Printf("Compare %d to %d at position foo[%d]\n",bar,foo[j],j)
            fmt.Println("foo: ",foo)
            fmt.Printf("foo[%d] = %d\n",j,foo[j])
            fmt.Println("foo: ",foo)
            fmt.Printf("foo[0:%d] = %d\n",j,foo[0:j])
            fmt.Printf("foo[:%d] = %d\n",j,foo[:j])
            fmt.Printf("foo[1:1] = %d\n",foo[1:1])

            fmt.Println("foo: ",foo)
            fmt.Println("bar: ",bar)

            //This line breaks stuff (WTF)
            fmt.Printf("append(foo[0:%d],bar) = %d\n",j,append(foo[0:j],bar))
            fmt.Println("foo: ",foo)
            fmt.Println("foo prior to assigning front half: ",foo)
            fmt.Println("foo: ",foo)
            frontHalf := append(foo[0:j],bar)
            fmt.Println("foo: ",foo)
            fmt.Printf("foo[%d:len(foo)] = %d\n",j,foo[j:len(foo)])
            backHalf := foo[j:len(foo)]

            fmt.Println("FrontHalf: ",frontHalf)
            fmt.Println("BackHalf: ",backHalf)
            ret = append(frontHalf,backHalf...)
            break
        }
        j++
    }

    fmt.Println("Output: ",ret)
}