package main

import (
    "fmt"
    "math"
    "math/big"
    "strconv"
)

func main() {
    // Assuming that each digit is random (ie. avg to 5), so the sum of the digits of
    // 2^1000, should be roughly 5*301 or ~1500

    // Great! It works against our test case!
    fmt.Println(niavePowerSum(2, 15))

    // Uh oh. I have a feeling that we hit an int max somewhere.
    fmt.Println(niavePowerSum(2, 1000))

    // Testing against 100, we see we get the same answer, 89, we must've hit some limit
    fmt.Println(niavePowerSum(2, 100))

    // New Implementation! This one converts to `math.big.Int`s to prevent overflowing.
    // Let's check against our normal case.
    fmt.Println(bigPowerSum(2, 15))

    // And here we are: 1366, same order of magnitude that we estimated.
    fmt.Println(bigPowerSum(2, 1000))

}

func niavePowerSum(n float64, k float64) int {
    power := math.Pow(n, k); // This is probably our culprit! Damn `int`s!
    powerString := strconv.Itoa(int(power))
    sum := 0;
    // This part looks good. No issue should arrise if we're going char by char.
    for i := 0; i < len(powerString); i++ {
        temp, _ := strconv.Atoi(string(powerString[i]))

        sum += temp
    }
    return sum
}

func bigPowerSum(n int, k int) int {
    var product big.Int

    product.SetUint64(uint64(n))
    multi := product

    for i := 0; i < k - 1; i++ {
        product.Mul(&product, &multi)
    }

    powerString := product.Text(10)
    sum := 0;

    for i := 0; i < len(powerString); i++ {
        temp, _ := strconv.Atoi(string(powerString[i]))

        sum += temp
    }
    return sum
}
