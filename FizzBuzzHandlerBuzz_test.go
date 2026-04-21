package main

import "testing"

func TestFizzBuzzHandlerFunctionBuzz(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{name: "four", input: 4, expected: "4"},
		{name: "five", input: 5, expected: "Buzz"},
		{name: "ten", input: 10, expected: "Buzz"},
		{name: "twenty", input: 20, expected: "Buzz"},
		{name: "twentyfive", input: 25, expected: "Buzz"},
		{name: "forty", input: 40, expected: "Buzz"},
		{name: "fifty", input: 50, expected: "Buzz"},
		{name: "seventy", input: 70, expected: "Buzz"},
		{name: "eighty", input: 80, expected: "Buzz"},
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
