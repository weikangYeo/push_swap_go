package algo

import (
	"fmt"
	"push_swap/model"
	"push_swap/stack"
	"push_swap/utils"
	"slices"
)

type simulateStep struct {
	raCount  int
	rraCount int
	rbCount  int
	rrbCount int
	rrCount  int
	rrrCount int
}

func (step simulateStep) totalCost() int {
	return step.raCount + step.rbCount + step.rbCount + step.rrbCount + step.rrCount + step.rrrCount
}

// try to find the longest sorted number and track the index
func RunLongIncrementAlgo(inputSlice []int) {
	bestSequence := getBestSequence(inputSlice)
	fmt.Printf("DEBUG - Best seq : %#v\n", bestSequence)
	stackA := stack.CreateStack(inputSlice)

	// Best Case
	if stackA.IsSorted() {
		return
	}

	var stackB *model.Node
	stackA, stackB = pushNotLisToStackB(bestSequence, inputSlice, stackA, stackB)

	bSize := stackB.Length()
	for range bSize {
		stackA, stackB = smartInsertBToA(stackA, stackB)
	}

}

func smartInsertBToA(stackA *model.Node, stackB *model.Node) (*model.Node, *model.Node) {
	// calculate step algorithm
	// find the least rotate step before do pushB
	// value will count from StackB pov
	bArrayValues := stackB.GetArrValue()
	bSize := stackB.Length()
	aArrayValues := stackA.GetArrValue()
	aSize := stackA.Length()

	var bestStep *simulateStep

	// loop current state of B
	for i := range bSize {
		bCurrentValue := bArrayValues[i]
		idx := getIndexToInsert(bCurrentValue, aArrayValues)
		proposedSteps := calculateRequiredSteps(idx, aSize, i, bSize)

		if bestStep == nil || proposedSteps.totalCost() < bestStep.totalCost() {
			bestStep = &proposedSteps
		}
	}

	return executeBestSteps(stackA, stackB, *bestStep)

}

func executeBestSteps(stackA *model.Node, stackB *model.Node, bestStep simulateStep) (*model.Node, *model.Node) {

	for range bestStep.raCount {
		stackA = stack.RotateA(stackA)
	}

	for range bestStep.rraCount {
		stackA = stack.ReverseRotateA(stackA)
	}

	for range bestStep.rbCount {
		stackB = stack.RotateB(stackB)
	}

	for range bestStep.rrbCount {
		stackB = stack.ReverseRotateB(stackB)
	}

	for range bestStep.raCount {
		stackA, stackB = stack.RotateBoth(stackA, stackB)
	}

	for range bestStep.rraCount {
		stackA, stackB = stack.ReverseRotateBoth(stackA, stackB)
	}

	return stackA, stackB
}

func calculateRequiredSteps(targetIdxInA int, sizeOfA int, bCurrentIdx int, sizeOfB int) simulateStep {

	steps := simulateStep{}
	halfIdx := sizeOfA / 2

	// if it is first half
	if targetIdxInA <= halfIdx {
		steps.raCount = targetIdxInA
	} else {
		steps.rraCount = targetIdxInA - halfIdx + 1
	}

	// how many step need to rb/rrb before we can get this B element
	bHalfIdx := sizeOfB / 2
	// if it is first half
	if bCurrentIdx <= bHalfIdx {
		steps.rbCount = bCurrentIdx
	} else {
		steps.rrbCount = bCurrentIdx - bHalfIdx + 1
	}

	// try to consolidate ra & rb to rr, rra & rrb to rrr
	if steps.raCount > 0 && steps.rbCount > 0 {
		steps.rrCount = utils.GetPositiveDiff(steps.raCount, steps.rbCount)
		steps.raCount = utils.GetPositiveDiff(steps.raCount, steps.rrCount)
		steps.rbCount = utils.GetPositiveDiff(steps.rbCount, steps.rrCount)
	}

	if steps.rraCount > 0 && steps.rrbCount > 0 {
		steps.rrCount = utils.GetPositiveDiff(steps.rraCount, steps.rrbCount)
		steps.rraCount = utils.GetPositiveDiff(steps.rraCount, steps.rrCount)
		steps.rrbCount = utils.GetPositiveDiff(steps.rrbCount, steps.rrCount)
	}

	return steps
}

func getIndexToInsert(valueToInsert int, relativeSortedArr []int) int {

	// array is empty, caller can just insert to idx 0
	if len(relativeSortedArr) == 0 {
		return 0
	}

	smallestValue, largestValue := relativeSortedArr[0], relativeSortedArr[0]
	var idxToInsert, smallestIdx, largestIdx int

	// not a good reading code in my pov,
	// but i want to use 1 loop to get 3 metrics
	for i := len(relativeSortedArr) - 1; i >= 0; i-- {
		if relativeSortedArr[i] < smallestValue {
			smallestValue = relativeSortedArr[i]
			smallestIdx = i
		}

		if relativeSortedArr[i] > largestValue {
			largestValue = relativeSortedArr[i]
			largestIdx = i
		}

		if valueToInsert < relativeSortedArr[i] {
			idxToInsert = i
		}
	}

	if valueToInsert < smallestValue {
		return smallestIdx
	}

	if valueToInsert > largestIdx {
		return largestIdx
	}

	return idxToInsert

}

func pushNotLisToStackB(bestSequence []int, inputSlice []int, stackA *model.Node, stackB *model.Node) (*model.Node, *model.Node) {
	countToPush := stackA.Length() - len(bestSequence)
	pushCount := 0

	// when pushCount reached, mean remaining item in stackA no need further RA (to save step)
	for pushCount < countToPush {
		if !slices.Contains(bestSequence, stackA.Value) {
			stackA, stackB = stack.PushToB(stackA, stackB)
			pushCount++
		} else {
			stackA = stack.RotateA(stackA)
		}
	}

	return stackA, stackB
}

func getBestSequence(inputSlice []int) []int {
	// Act as first Node which value that are cheaper to sort everything
	bestRelativeSortedCount := -1
	var bestSortedSequence []int
	// iterate inputSlice, see which starting idx get more sorted element
	for idx := range inputSlice {
		previousValue := inputSlice[idx]
		relativeSortedCount := 0
		var sortedSequence []int

		// loop from start idx to end of slice
		for i := idx + 1; i < len(inputSlice); i++ {
			if inputSlice[i] > previousValue {
				relativeSortedCount++
				sortedSequence = append(sortedSequence, inputSlice[i])
			}

			previousValue = inputSlice[i]
		}

		// loop remaining items from start of slice to start of idx
		// consider it is a looped list, need to iterate back to the idx
		// e.g. 1,2,3,4 if start with "3", after "4", need to continue to "1" & "2"
		for i := 0; i < idx; i++ {
			if inputSlice[i] > previousValue {
				relativeSortedCount++
				sortedSequence = append(sortedSequence, inputSlice[i])
			}
			previousValue = inputSlice[i]
		}

		

		if relativeSortedCount > bestRelativeSortedCount {
			bestRelativeSortedCount = relativeSortedCount
			bestSortedSequence = sortedSequence
		}
	}

	return bestSortedSequence
}
