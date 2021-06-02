package main

import "fmt"

// Multiples of 3 and 5
func main() {
    sum := 0
    max := 1000

    for i := 0; i < max; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            sum += i
        }
    }

    fmt.Println(sum)
}
