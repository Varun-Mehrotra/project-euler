package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)


func main() {
	numbersFile, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer numbersFile.Close()

	numberPyramidList := make([][]int, 0)

	scanner := bufio.NewScanner(numbersFile)
	for scanner.Scan() {
		lineNumbers := scanner.Text()

		numbers := strings.Split(lineNumbers, " ")

		intNumbers := make([]int, 0)
		
		for _, number := range numbers {
			intNumber, _ := strconv.Atoi(number)
			intNumbers = append(intNumbers, intNumber )
		}
		
		numberPyramidList = append(numberPyramidList, intNumbers)
	}



	for i := len(numberPyramidList)-2; i >= 0; i-- {
		for j, number := range numberPyramidList[i] {
			fmt.Println(i, j, number)
			rightNumber := numberPyramidList[i+1][j]
			leftNumber := numberPyramidList[i+1][j+1]

			if rightNumber > leftNumber {
				numberPyramidList[i][j] += rightNumber
			} else {
				numberPyramidList[i][j] += leftNumber
			}
		}
	}

	fmt.Println(numberPyramidList)
}

