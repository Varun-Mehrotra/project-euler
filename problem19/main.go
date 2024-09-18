package main

import (
	"fmt"
)


func main() {

	done := false
	dayOfWeek := 1
	year := 1900
	dayOfMonth := 1
	month := 1
	isLeapYear := false

	monthDays := map[int]int{
        1: 31,
        2: 28,
        3: 31,
		4: 30,
        5: 31,
        6: 30,
		7: 31,
        8: 31,
		9: 30,
        10: 31,
        11: 30,
        12: 31,
	}

	firstMonthSunday := 0
	// This loop represents incrementing a day
	for !done {
		isLeapYear = false 
		if year % 100 == 0{
			if year % 400 == 0 {
				isLeapYear = true
			}
		} else if year % 4 == 0 {
			isLeapYear = true 
		}

		daysInThisMonth := monthDays[month]

		if isLeapYear && month == 2 {
			daysInThisMonth = 29
		}

		if year >= 1901 && dayOfWeek == 7 && dayOfMonth == 1 {
			firstMonthSunday++
		}
		fmt.Println(year, month, dayOfMonth, dayOfWeek)


		if year == 2000 && month == 12 && dayOfMonth > 2 {
			done = true
		}

		// Increment day of the week every day
		dayOfWeek = incrementDayOfWeek(dayOfWeek)

		if dayOfMonth == daysInThisMonth && month == 12 {
			month = 1
			dayOfMonth = 0
			year++
		} else if dayOfMonth == daysInThisMonth {

			month++
			dayOfMonth = 0
		}
		dayOfMonth++

	}

	fmt.Println(firstMonthSunday)
}

func incrementDayOfWeek(oldDay int) int {
	if oldDay + 1 >= 8 {
		return 1
	}
	return oldDay+1
}

