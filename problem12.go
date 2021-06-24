package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(getTriangleNumberWithFactors(500))
}

func calcTriangeNumbers(n int) int {
    return n*(n+1)/2
}

func getNFactors(n int) int {
    factorsCount := 0;
    // Only go up to centeroid of factors
    for i := 1; i < int(math.Sqrt(float64(n))); i++ {
        if n % i == 0 {
            factorsCount += 1
        }
    }
    // Double factors cause we only went to centeroid
    return factorsCount*2
}

func getTriangleNumberWithFactors(factors int) int{
    i := 0
    for {
        triangleNum := calcTriangeNumbers(i)
        if getNFactors(triangleNum) > factors {
            return triangleNum
        }
        i++
    }
    return 0
}
