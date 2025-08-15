package main

import (
	"fmt"
	"os"
	"push_swap/algo"
	"push_swap/sort"
	"push_swap/utils"
)

func main() {
	intSlice, err := utils.ReadIntSliceFromTerminal()
	if err != nil {
		// use os.Stderr so it will write to stderr in terminal, which most of open source checker is check based on this
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if intSlice == nil {
		// print nothing and stop program as requirement asked
		return
	}

	sortedSlice := sort.MergeSort(intSlice)

	algo.RunSimpleAlgo(intSlice, sortedSlice)
}
