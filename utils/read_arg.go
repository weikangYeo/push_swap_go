package utils

import (
	"errors"
	"math"
	"os"
	"strconv"
)

// ParseInts converts a slice of strings to a slice of ints.
// Spin off so the core logic could be tested via unit test
func ParseInts(args []string) ([]int, error) {
	if len(args) == 0 {
		// Return a non-nil empty slice for consistency.
		return []int{}, nil
	}

	// Create a slice of the correct size.
	numbers := make([]int, len(args))

	numMap := make(map[int]int)
	for i, argStr := range args {
		num, err := strconv.Atoi(argStr)
		if err != nil {
			// invalid input found, but show generic Error as requirement asked
			return nil, errors.New("Error")
		}

		if num < math.MinInt32 || num > math.MaxInt32 {
			// show generic Error as requirement asked
			return nil, errors.New("Error")
		}

		_, isExists := numMap[num]
		if isExists {
			// show generic Error as requirement asked
			return nil, errors.New("Error")
		} else {
			numMap[num] = 1
		}
		numbers[i] = num
	}
	return numbers, nil
}

// ReadIntSliceFromTerminal is a wrapper around ParseInts that uses os.Args.
func ReadIntSliceFromTerminal() ([]int, error) {
	return ParseInts(os.Args[1:])
}
