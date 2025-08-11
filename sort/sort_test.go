package sort

import (
	"reflect"
	"testing"
)

func TestSortSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "single element",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "sorted slice",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "reverse sorted slice",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "unsorted slice with duplicates",
			input: []int{3, 1, 4, 1, 5, 9, 2, 6},
			want:  []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
		{
			name:  "slice with negative numbers",
			input: []int{-3, 1, -4, 0, 5, -9},
			want:  []int{-9, -4, -3, 0, 1, 5},
		},
        {
			name:  "slice with all same elements",
			input: []int{2, 2, 2, 2, 2},
			want:  []int{2, 2, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSort(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
