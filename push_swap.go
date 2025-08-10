package main

import (
	"push_swap/model"
	"push_swap/utils"
	"fmt"
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

}
