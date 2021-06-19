package main

import (
    "fmt"
)

//Smallest multiple
func main() {
    fmt.Println(smallestMultiple(10))
}

func smallestMultiple(n int) int {
    i := n
    for {
        if evenDivisible(i, 20) {
            return i
        }
        i++
    }
    return 0
}

func evenDivisible(n int, multipleMax int) bool {
    for i := multipleMax/2; i < multipleMax; i++ {
        if n % i != 0 {
            return false
        }
    }
    return true
}
