package main

import (
    "fmt"
    "math"
)

//10001st Prime
func main() {
    n := 10001
    fmt.Println(findNthPrime(n))
}

func findNthPrime(n int) int {
    primes := []int{}
    i := 1
    for {
        if isPrime(i) {
            primes = append(primes, i)
            if len(primes) == n {
                break
            }
        }
        i++
    }
    return primes[len(primes) - 1]
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
