package main

import "fmt"

func main() {
    num := 600851475143
    fmt.Println(recursivePrimeFactors(num))
}

func recursivePrimeFactors(n int) []int {
    primeFactors := []int{}
    for i := 2; i < n; i++ {
        if (n % i == 0) {
            return append(recursivePrimeFactors(n / i), i)
        }
    }
    return append(primeFactors, n)
}
