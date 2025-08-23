package utils

// Get number of different between 2 number, value will always return POSITIVE number (>=0) value
func GetPositiveDiff(a, b int) int {

	result := a - b
	if result < 0 {
		result = result * -1
	}

	return result
}

func GetMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}