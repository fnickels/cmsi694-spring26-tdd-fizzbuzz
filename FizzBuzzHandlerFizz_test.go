package main

import "testing"

func TestFizzBuzzHandlerFunctionFizz(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{name: "one", input: 1, expected: "1"},
		{name: "two", input: 2, expected: "2"},
		{name: "three", input: 3, expected: "Fizz"},
		{name: "six", input: 6, expected: "Fizz"},
		{name: "nine", input: 9, expected: "Fizz"},
		{name: "ninetythree", input: 93, expected: "Fizz"},
		{name: "ninetysix", input: 96, expected: "Fizz"},
		{name: "ninetynine", input: 99, expected: "Fizz"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := FizzBuzzHandler(tc.input)
			if actual != tc.expected {
				t.Fatalf("FizzBuzzHandler(%d) = %q, want %q", tc.input, actual, tc.expected)
			}
		})
	}
}
