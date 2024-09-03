package main

import "fmt"

// Even Fibonacci numbers
func main() {
    max := 4000000
    prev := 1
    sum := 0
    prev_prev := 0
    i := 1
    for i < max {
        prev_prev = prev
        prev = i
        i = prev + prev_prev
        if i % 2 == 0 {
            sum += i
        }
    }

    fmt.Println(sum)
}

// -- IGNORE ALL THIS, just a concept if fibonacciSequence is needed again.

// func fib(n int) int {
//     if n == 1 {
//         return 1
//     }
//     if n == 0 {
//         return 0
//     }
//     return fib(n-1) + fib(n-2)
// }

// type fibonacciSequence struct {
//     sequence []int
//     currentVal int
//     currentN int
//     val int
// }
//
//
// func newFibonacciSequence(n int) fibonacciSequence {
//     seq := []int{1, 1}
//     seq.append(seq, 1)
//
//
//     p := fibonacciSequence{sequence: seq }
// }
//
// func (fibonacciSequence) getSequence(n int) []int {
//     for i := 0; i < n; n
// }
//
// func nextInSequence() {
//     someth
//
// }
