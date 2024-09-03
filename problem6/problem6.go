package main

import (
    "fmt"
)

//Sum square difference
func main() {
    n := 100
    fmt.Println(squareOfSums(n) - sumOfSquares(n))
}

func sumOfSquares(n int) int {
    sum := 0
    for i := 0; i <= n; i++ {
        sum += i*i
    }
    return sum
}

func squareOfSums(n int) int {
    sum := 0
    for i := 0; i <= n; i++ {
        sum += i
    }
    return sum*sum
}
