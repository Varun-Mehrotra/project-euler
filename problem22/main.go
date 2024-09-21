package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)


func main() {	
	namesFile, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer namesFile.Close()

	names := make([]string, 0)

	scanner := bufio.NewScanner(namesFile)
	for scanner.Scan() {
		namesStringList := strings.Split(scanner.Text(), ",")

		for _, name := range namesStringList {
			names = append(names, name)
		}
		
	}

	sort.Strings(names)

	// names := []string{"ZZZ"}

	runeMap := make(map[rune]int)

    // Fill the map with values
    for i := 'A'; i <= 'Z'; i++ {
        runeMap[i] = int(i - 'A' + 1)
    }

	totalNameScore := 0
	
	for i, name := range names {
		nameSum := 0
		for _, letter := range name {
			nameSum += runeMap[letter]
		}

		totalNameScore += nameSum * (i+1)
	}

	fmt.Println(totalNameScore)
}