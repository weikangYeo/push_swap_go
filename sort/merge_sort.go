package sort

func MergeSort(intSlice []int) []int {
	if len(intSlice) < 2 {
		return intSlice
	}

	half := len(intSlice) / 2
	return merge(intSlice[:half], intSlice[half:])
}

func merge(left, right []int) []int {

	if len(left) > 1 {
		half := len(left) / 2
		left = merge(left[:half], left[half:])
	}

	if len(right) > 1 {
		half := len(right) / 2
		right = merge(right[:half], right[half:])
	}

	sizeOfSlice := len(left) + len(right)
	result := make([]int, sizeOfSlice)
	var leftPtr, rightPtr, i int

	for i < sizeOfSlice && leftPtr < len(left) && rightPtr < len(right) {
		if left[leftPtr] < right[rightPtr] {
			result[i] = left[leftPtr]
			leftPtr++
		} else {
			result[i] = right[rightPtr]
			rightPtr++
		}
		i++
	}

	// fill in remaining number
	for leftPtr < len(left) {
		result[i] = left[leftPtr]
		leftPtr++
		i++
	}

	for rightPtr < len(right) {
		result[i] = right[rightPtr]
		rightPtr++
		i++
	}

	return result
}
