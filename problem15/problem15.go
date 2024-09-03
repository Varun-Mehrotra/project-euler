package main

import "fmt"

func main() {
    fmt.Println(searchGridPascal(20))
}

func searchGrid(x int, y int) int {
    if x == 0 || y == 0 {
        return 1
    }
    return searchGrid(x - 1, y) + searchGrid(x, y - 1)
}

func searchGridPascal(n int) int {
    c := 1;
    for i := 1; i <= n; i++ {
        c = c * (n + i) / i
    }
    return c
}
