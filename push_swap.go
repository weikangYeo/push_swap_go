package main

import (
	"fmt"
	"push_swap/algo"
	"push_swap/sort"
	"push_swap/utils"
)

func main() {
	intSlice, err := utils.ReadIntSliceFromTerminal()
	if err != nil {
		// any input error will just print out and stop program.
		fmt.Print(err)
		return
	}
	if intSlice == nil {
		// print nothing and stop program as requirement asked
		return
	}

	sortedSlice := sort.MergeSort(intSlice)

	algo.RunSimpleAlgo(intSlice, sortedSlice)
}
