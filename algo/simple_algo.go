package algo

import (
	"errors"
	"push_swap/model"
	"push_swap/stack"
)

// simple algo is simply just iterate the stack A, and find the smallest,
// then push to stack B, when A is empty push it back A
func RunSimpleAlgo(inputSlice, sortedSlice []int) error {
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
