package main

import "strconv"

func FizzBuzzHandler(value int) string {

	switch value % 15 {
	case 0:
		return "FizzBuzz"

	case 3:
		fallthrough
	case 6:
		fallthrough
	case 9:
		fallthrough
	case 12:
		return "Fizz"

	case 5:
		fallthrough
	case 10:
		return "Buzz"
	}

	return strconv.Itoa(value)
}
