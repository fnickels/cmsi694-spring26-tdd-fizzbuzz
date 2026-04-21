# cmsi694-spring26-tdd-fizzbuzz
TDD Demo Framework


## Requirements
Given: an input of numbers from 1–100
When:
* A number is a multiple of ‘3’ return “Fizz”
* A number is a of ‘5’ return “Buzz”
* A number is a of both ‘3’ and ‘5’ return “FizzBuzz”
* A number is not divisible by ‘3’ or ‘5’ return the number itself

## Expected output: 

```
   1, 2, Fizz, 4, Buzz, ……, 14, FizzBuzz, 16, …
```

## Running the application

From the project root, execute:

```bash
go run .
```

This prints each value from 1 through 100 on its own line.

## Running tests

Run all tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```
