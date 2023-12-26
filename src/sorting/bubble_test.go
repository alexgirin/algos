package sorting

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
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
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			BubbleSort[int](&test.input)
			assert.True(t, reflect.DeepEqual(&test.input, &test.result))
		})
	}
}
