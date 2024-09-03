package main

import (
    "strconv"
    "fmt"
)
// Largest Palindrome Product
func main() {
    fmt.Println(largestPalindromeProduct(999))
}

func largestPalindromeProduct(max int) int {
    biggestPalindrome := 0
    lowerJ := 0
    for i := max; i > 1; i-- {
        for j := max; j > lowerJ; j-- {
            fmt.Println("i", i, ", j", j)
            if isPalindrome(strconv.Itoa(i*j)) && i*j > biggestPalindrome {
                biggestPalindrome = i*j
                break
            }
        }
        lowerJ = biggestPalindrome / (i - 1)
    }
    return biggestPalindrome
}

func isPalindrome(str string) bool {
    for i := 0; i < len(str) / 2; i++ {
        if str[i] != str[len(str)-i - 1] {
            return false
        }
    }
    return true
}
