package main

import (
    "fmt"
)

// Pythagorean Triplet
func main() {
    fmt.Println(findPythagoreanTriplets(1000))
}

func calcA(m int, n int) int {
    return m*m - n*n
}

func calcB(m int, n int) int {
    return 2*m*n
}

func calcC(m int, n int) int {
    return m*m + n*n
}

func findPythagoreanTriplets(sum int) int {
    a := 0
    b := 0
    c := 0
    n := 1
    m := n + 1
    for {
        a = calcA(m, n)
        b = calcB(m, n)
        c = calcC(m, n)

        if a + b + c == sum {
            return a*b*c
        }

        if a + b + c > sum {
            n++
            m = n
        }
        m++
    }
    return 0
}
