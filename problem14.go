package main

import "fmt"

var helperSieve map[int]int // This would've been cool, but unrequired.

func main() {
    maxVal := 0;
    maxInt := 0;
    curVal := 0;

    for i := 2; i < 1000000; i++ {
        if i % 2 == 0 {
            curVal = nEven(i);
        } else {
            curVal = nOdd(i);
        }
        if curVal > maxVal {
            maxVal = curVal
            maxInt = i;
        }
    }
    fmt.Println(maxInt)
}

func nEven(n int) int {
    val := n / 2
    if val == 1 {
        return 2
    } else if val % 2 == 0 {
        return nEven(val) + 1
    }
    return nOdd(val) + 1
}

func nOdd(n int) int {
    val := 3 * n + 1
    if val == 1 {
        return 2
    } else if val % 2 == 0 {
        return nEven(val) + 1
    }
    return nOdd(val) + 1
}
