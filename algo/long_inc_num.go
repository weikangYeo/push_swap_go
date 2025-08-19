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
	if inputSlice == nil {
		return
	}

	bestSequence := getBestSequence(inputSlice)
	stackA := stack.CreateStack(inputSlice)

	fmt.Printf("best Seq: %#v\n", bestSequence)

	// Best Case
	if stackA.IsSorted() {
		return
	}

	var stackB *model.Node
	stackA, stackB = pushNotLisToStackB(bestSequence, stackA, stackB)

	fmt.Printf("After Push Non LIS to B\n")
	fmt.Printf("stackA: %#v\n", stackA.ToSlice())
	fmt.Printf("stackB: %#v\n", stackB.ToSlice())

	bSize := stackB.Length()
	for range bSize {
		stackA, stackB = smartInsertBToA(stackA, stackB)
	}

	fmt.Printf("Before Final Rotation !\n")
	fmt.Printf("stackA: %#v\n", stackA.ToSlice())
	fmt.Printf("stackB: %#v\n", stackB.ToSlice())

	// final rotation
	relativedSortedA := stackA.ToSlice()

	// stackA would never have empty element case, as the caller is already handled it before call
	smallestIdx := 0
	smallestValue := relativedSortedA[0]
	for i, value := range relativedSortedA {
		if value < smallestValue {
			smallestValue = value
			smallestIdx = i
		}
	}

	// first half
	if smallestIdx <= stackA.Length()/2 {
		for range smallestIdx {
			stackA = stack.RotateA(stackA)
		}
	} else {
		for range stackA.Length() - smallestIdx {
			stackA = stack.ReverseRotateA(stackA)
		}
	}

	// fmt.Printf("Final !\n")
	// fmt.Printf("stackA: %#v\n", stackA.ToSlice())
	// fmt.Printf("stackB: %#v\n", stackB.ToSlice())

}

func smartInsertBToA(stackA *model.Node, stackB *model.Node) (*model.Node, *model.Node) {
	// calculate step algorithm
	// find the least rotate step before do pushB
	// value will count from StackB pov
	bValues := stackB.ToSlice()
	bSize := stackB.Length()
	aValues := stackA.ToSlice()
	aSize := stackA.Length()

	var bestStep *simulateStep

	// loop current state of B
	for i := range bSize {
		bCurrentValue := bValues[i]
		idx := getIndexToInsert(bCurrentValue, aValues)
		proposedSteps := calculateRequiredSteps(idx, aSize, i, bSize)

		if bestStep == nil || proposedSteps.totalCost() < bestStep.totalCost() {
			// fmt.Printf("Current IDX in B %d, B value: %d, idx to insert in A: %d\n", i, bCurrentValue, idx)
			bestStep = &proposedSteps
		}
	}

	return executeBestSteps(stackA, stackB, *bestStep)

}

func executeBestSteps(stackA *model.Node, stackB *model.Node, bestStep simulateStep) (*model.Node, *model.Node) {

	// fmt.Printf("Plan: %#v\n", bestStep)

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

	for range bestStep.rrCount {
		stackA, stackB = stack.RotateBoth(stackA, stackB)
	}

	for range bestStep.rrrCount {
		stackA, stackB = stack.ReverseRotateBoth(stackA, stackB)
	}

	// after all the rotation, push the value from B to A
	stackA, stackB = stack.PushToA(stackA, stackB)

	return stackA, stackB
}

func calculateRequiredSteps(targetIdxInA int, sizeOfA int, bCurrentIdx int, sizeOfB int) simulateStep {

	steps := simulateStep{}

	// targetIdxInA == sizeOfA and targetIdxInA == 0 are best case,
	// can directly do pa
	// for targetIdxInA == sizeOfA mean it is suppose to enter at end of stack A
	// can directly pa because it is still logical sorted.
	if targetIdxInA != 0 && targetIdxInA != sizeOfA {
		// first half
		if targetIdxInA <= sizeOfA/2 {
			steps.raCount = targetIdxInA
		} else {
			steps.rraCount = sizeOfA - targetIdxInA
		}
	}

	// how many step need to rb/rrb before we can get this B element
	if bCurrentIdx <= sizeOfB/2 {
		steps.rbCount = bCurrentIdx
	} else {
		steps.rrbCount = sizeOfB - bCurrentIdx
	}

	// try to consolidate ra & rb to rr, rra & rrb to rrr
	if steps.raCount > 0 && steps.rbCount > 0 {
		steps.rrCount = utils.GetPositiveDiff(steps.raCount, steps.rbCount)
		steps.raCount = utils.GetPositiveDiff(steps.raCount, steps.rrCount)
		steps.rbCount = utils.GetPositiveDiff(steps.rbCount, steps.rrCount)
	}

	if steps.rraCount > 0 && steps.rrbCount > 0 {
		steps.rrrCount = utils.GetPositiveDiff(steps.rraCount, steps.rrbCount)
		steps.rraCount = utils.GetPositiveDiff(steps.rraCount, steps.rrCount)
		steps.rrbCount = utils.GetPositiveDiff(steps.rrbCount, steps.rrCount)
	}

	return steps
}

// return the idx to insert valueToInsert.
// Given getIndexToInsert(1, [2, 4, 6, 8]) return 0,
// mean new proposed sortedSlice = [1, 2, 4, 6, 8].
// getIndexToInsert(9, [2, 4, 6, 8]), return 4
// getIndexToInsert(1, [6, 8, 2 ,4]), return 2
func getIndexToInsert(valueToInsert int, logicalSortedSlice []int) int {

	// array is empty, caller can just insert to idx 0
	if len(logicalSortedSlice) == 0 {
		return 0
	}

	smallestValue, largestValue := logicalSortedSlice[0], logicalSortedSlice[0]
	var smallestIdx, largestIdx int

	for i, value := range logicalSortedSlice {
		if value < smallestValue {
			smallestValue = value
			smallestIdx = i
		}

		if value > largestValue {
			largestValue = value
			largestIdx = i
		}
	}

	if valueToInsert < smallestValue {
		return smallestIdx
	}

	if valueToInsert > largestValue {
		return largestIdx + 1
	}

	// the following 2 loop it to find out where should the valueToInsert to insert
	// using 2 loop is bcoz is to start from first logical sorted number (smallest idx)
	// then use second loop to continue until the last logical sorted number (largest idx)
	for i := smallestIdx; i < len(logicalSortedSlice); i++ {
		if logicalSortedSlice[i] > valueToInsert {
			return i
		}
	}
	for i := 0; i <= largestIdx; i++ {
		if logicalSortedSlice[i] > valueToInsert {
			return i
		}
	}

	// runtime logical error here
	return -1
}

func pushNotLisToStackB(bestSequence []int, stackA *model.Node, stackB *model.Node) (*model.Node, *model.Node) {
	countToPush := stackA.Length() - len(bestSequence)
	pushCount := 0

	// fmt.Printf("countToPush %d\n", countToPush)

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
		sortedSequence = append(sortedSequence, previousValue)

		// loop from start idx to end of slice
		for i := idx + 1; i < len(inputSlice); i++ {
			if inputSlice[i] > previousValue {
				relativeSortedCount++
				sortedSequence = append(sortedSequence, inputSlice[i])
				previousValue = inputSlice[i]
			}
		}

		// loop remaining items from start of slice to start of idx
		// consider it is a looped list, need to iterate back to the idx
		// e.g. 1,2,3,4 if start with "3", after "4", need to continue to "1" & "2"
		for i := 0; i < idx; i++ {
			if inputSlice[i] > previousValue {
				relativeSortedCount++
				sortedSequence = append(sortedSequence, inputSlice[i])
				previousValue = inputSlice[i]
			}
		}

		if relativeSortedCount > bestRelativeSortedCount {
			bestRelativeSortedCount = relativeSortedCount
			bestSortedSequence = sortedSequence
		}
	}

	return bestSortedSequence
}
