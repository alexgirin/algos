package sorting

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		title  string
		input  []int
		result []int
	}{
		{
			title:  "Empty array",
			input:  make([]int, 0),
			result: make([]int, 0),
		},
		{
			title:  "One element",
			input:  []int{1},
			result: []int{1},
		},
		{
			title:  "Many elements",
			input:  []int{13, 1, 8, 2, 5, 3, 1},
			result: []int{1, 1, 2, 3, 5, 8, 13},
		},
		{
			title:  "Negative/positive elements",
			input:  []int{13, -1, 8, -2, -5, 3, 1},
			result: []int{-5, -2, -1, 1, 3, 8, 13},
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result := MergeSort[int](&test.input)
			assert.True(t, reflect.DeepEqual(&result, &test.result))
		})
	}
}
