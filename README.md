# PUSH SWAP

Just a side project to grill my Go lang skill, get me out from comfort zone (since I am not a Go person)

## How to run
### Run via Go Run
`go run push_swap.go 53 22 17 4 56 2 65 97 1 86 64 83 69 96 73 41 0 99 3 75 26 82 9 23 12 78 33 68 95 63 89 47 77 42 35 58 79 50 5 36 70 76 7 92 61 11 51 30 60 98 25 72 46 31 81 91 16 55 84 52 93 29 32 14 67 74 40 34 44 49 18 38 20 59 6 39 48 37 8 27 15 10 24 54 43 87 88 57 45 62 85 28 94 80 13 71 66 19 90 21`

### Run via Executable
```sh
go build  
```
```sh
./push_swap 53 22 17 4 56 2 65 97 1 86 64 83 69 96 73 41 0 99 3 75 26 82 9 23 12 78 33 68 95 63 89 47 77 42 35 58 79 50 5 36 70 76 7 92 61 11 51 30 60 98 25 72 46 31 81 91 16 55 84 52 93 29 32 14 67 74 40 34 44 49 18 38 20 59 6 39 48 37 8 27 15 10 24 54 43 87 88 57 45 62 85 28 94 80 13 71 66 19 90 21 
```

### Run via checker
```sh
./push_swap 53 22 17 4 56 2 65 97 1 86 64 83 69 96 73 41 0 99 3 75 26 82 9 23 12 78 33 68 95 63 89 47 77 42 35 58 79 50 5 36 70 76 7 92 61 11 51 30 60 98 25 72 46 31 81 91 16 55 84 52 93 29 32 14 67 74 40 34 44 49 18 38 20 59 6 39 48 37 8 27 15 10 24 54 43 87 88 57 45 62 85 28 94 80 13 71 66 19 90 21 | ./checker_linux 53 22 17 4 56 2 65 97 1 86 64 83 69 96 73 41 0 99 3 75 26 82 9 23 12 78 33 68 95 63 89 47 77 42 35 58 79 50 5 36 70 76 7 92 61 11 51 30 60 98 25 72 46 31 81 91 16 55 84 52 93 29 32 14 67 74 40 34 44 49 18 38 20 59 6 39 48 37 8 27 15 10 24 54 43 87 88 57 45 62 85 28 94 80 13 71 66 19 90 21 
```

## How to unit test

This project includes unit tests to verify the correctness of the utility functions.

To run all tests, execute the following command from the root of the project:

```bash
go test ./...
```

This command will discover and run all test files (files ending in `_test.go`) in all sub-packages.

## Credit
Credit to https://github.com/gemartin99/Push-Swap-Tester script to validate executable functionality 