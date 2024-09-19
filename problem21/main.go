package main

import (
	"fmt"
	"math"
)


func main() {	
	amicableMap := make(map[int]int, 0)
	amicableSum := 0

	for i := 1; i < 10000; i++ {
		oldDofI, ok := amicableMap[i]

		dI := d(i)

		if ok && oldDofI == dI{
			fmt.Println("i", i)
			fmt.Println("d(i)", dI)

			amicableSum += i + dI
		}
		amicableMap[dI] = i
	}

	fmt.Println(amicableSum)
}

func d(n int) int {
	// find divisors
	properDivisorSum := 1;

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0{
			properDivisorSum += i + n/i
		}
	}

	// sum divisors
	return properDivisorSum
}