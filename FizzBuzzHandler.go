package main

import "strconv"

func FizzBuzzHandler(value int) string {

	switch value % 15 {
	case 0:
		return "FizzBuzz"

	case 3, 6, 9, 12:
		return "Fizz"

	case 5, 10:
		return "Buzz"
	}

	return strconv.Itoa(value)
}
