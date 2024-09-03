package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func main() {
	numbersFile, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer numbersFile.Close()

	sum := "0"

	scanner := bufio.NewScanner(numbersFile)
	for scanner.Scan() {
		number := scanner.Text()
		oldSum := sum

		sum = addBigNumbers(sum, number)

		fmt.Println(oldSum, " + ", number, " = ", sum)

	}
	fmt.Println("Sum:")
	fmt.Println(sum)

	// fmt.Println(addBigNumbers("0", "37107287533902102798797998220837590246510135740250"))
}

func addBigNumbers(x string, y string) string {
	sum := ""
	carryOver := 0

	var shorter string
	var longer string
	if len(x) > len(y) {
		shorter = y
		longer = x
	} else {
		shorter = x
		longer = y
	}
	// Goes through each digit and adds them 
	for i := 0; i < len(shorter); i++ {
		shorterDigit, _ := strconv.Atoi(string(shorter[len(shorter)-i-1]))
		longerDigit, _ := strconv.Atoi(string(longer[len(longer)-i-1]))
		digitSum :=  shorterDigit + longerDigit + carryOver

		if digitSum >= 10 {
			carryOver = 1
			digitSum = digitSum - 10
		} else {
			carryOver = 0
		}

		sum = strconv.Itoa(digitSum) + sum 
	}

	if len(x) != len(y) {
		leftOver := longer[0:len(longer)-len(shorter)]
		leftOver = addBigNumbers(leftOver, strconv.Itoa(carryOver))
		// fmt.Println("leftOver", leftOver)
		// fmt.Println("carryOver", carryOver)
		// fmt.Println("sum", sum)
		sum = leftOver + sum
	} else if carryOver > 0 {
		sum = strconv.Itoa(carryOver) + sum
	}

	return sum
}