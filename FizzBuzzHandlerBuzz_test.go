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
