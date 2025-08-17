package algo

import (
	"errors"
	"push_swap/model"
	"push_swap/sort"
	"push_swap/stack"
)

// simple algo is simply just iterate the stack A, and find the smallest,
// then push to stack B, when A is empty push it back A
// cons of this is will be too many steps, as it will fail out the min threshold and become KO directly
func RunSimpleAlgo(inputSlice []int) error {
	sortedSlice := sort.MergeSort(inputSlice)

	stackA := stack.CreateStack(inputSlice)
	var stackB *model.Node

	for _, currentValue := range sortedSlice {
		idx := stackA.IndexOf(currentValue)
		if idx < 0 {
			return errors.New("invalid idx")
		}

		sizeOfA := stackA.Length()
		// divide half to see it is first half or second half
		// then decide it is ra or rra to reach the value
		isFirstHalf := idx < (sizeOfA / 2)

		for stackA.Value != currentValue {
			if isFirstHalf {
				stackA = stack.RotateA(stackA)
			} else {
				stackA = stack.ReverseRotateA(stackA)
			}
		}
		stackA, stackB = stack.PushToB(stackA, stackB)
	}

	for stackB != nil {
		// push back to A
		stackA, stackB = stack.PushToA(stackA, stackB)

	}

	return nil
}
