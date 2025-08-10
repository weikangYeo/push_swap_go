package utils

import (
	"reflect" // Used for deep equality check of slices
	"testing"
)

func TestParseInts(t *testing.T) {
	// We use table-driven tests, which is a common pattern in Go.
	// It allows us to define a set of test cases easily.
	testCases := []struct {
		name      string // Name of the test case
		input     []string
		want      []int
		expectErr bool // Do we expect an error?
	}{
		{
			name:      "Happy path with positive numbers",
			input:     []string{"1", "2", "3", "100"},
			want:      []int{1, 2, 3, 100},
			expectErr: false,
		},
		{
			name:      "Empty input slice",
			input:     []string{},
			want:      []int{},
			expectErr: false,
		},
		{
			name:      "Nil input slice",
			input:     nil,
			want:      []int{},
			expectErr: false,
		},
		{
			name:      "Input with non-integer value",
			input:     []string{"1", "two", "3"},
			want:      nil,
			expectErr: true,
		},
		{
			name:      "Input with negative numbers",
			input:     []string{"-5", "0", "-10"},
			want:      []int{-5, 0, -10},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		// t.Run makes test results easier to read, grouping them by case name.
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseInts(tc.input)

			// Check if we got an error when we expected one, or vice-versa.
			hasErr := err != nil
			if hasErr != tc.expectErr {
				t.Fatalf("ParseInts() error = %v, expectErr %v", err, tc.expectErr)
			}

			// If we got the expected error state, we also need to check the result.
			// reflect.DeepEqual is used to compare slices and other complex types.
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ParseInts() = %v, want %v", got, tc.want)
			}
		})
	}
}
