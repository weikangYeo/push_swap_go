package algo

import "push_swap/stack"

// simple algo is simply just iterate the stack A, and find the smallest,
// then push to stack B, when A is empty push it back A
func RunSimpleAlgo(inputSlice, sortedSlice []int) {
	stackA := stack.CreateStack(inputSlice)
	stackB := nil

	for _, currentValue := range sortedSlice {
		idx := stackA.IndexOf(currentValue)
		
		// divide half to see it is first half or second half
		// then decide it is ra or rra to reach the value

	}
}
