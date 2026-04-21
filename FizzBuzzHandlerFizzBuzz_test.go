package main

import "testing"

func TestFizzBuzzHandlerFunctionFizzBuzz(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{name: "fourteen", input: 14, expected: "14"},
		{name: "fifteen", input: 15, expected: "FizzBuzz"},
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
