package main

import (
    "math"
    "fmt"
)

func main() {
    fmt.Println(isPrime(2))
    fmt.Println(sumPrimes(2000000))
}

func isPrime(n int) bool {
    if n == 1 {
        return false
    }

    for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}

func sumPrimes(n int) int {
    sum := 0
    for i := 0; i < n; i++ {
        if isPrime(i) {
            sum += i
        }
    }
    return sum
}
