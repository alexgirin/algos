package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearchInt(t *testing.T) {
	tests := []struct {
		title  string
		input  *TestInput[int]
		result int
	}{
		{
			"True",
			&TestInput[int]{
				haystack: []int{1, 2, 3, 4, 5},
				needle:   2,
			},
			1,
		},
		{
			"False",
			&TestInput[int]{
				haystack: []int{1, 2, 3, 4, 5},
				needle:   7,
			},
			-1,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result := Linear[int](&test.input.haystack, test.input.needle)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestLinearSearchString(t *testing.T) {
	tests := []struct {
		title  string
		input  *TestInput[string]
		result int
	}{
		{
			"True",
			&TestInput[string]{
				haystack: []string{"1", "2", "3", "4", "5"},
				needle:   "2",
			},
			1,
		},
		{
			"False",
			&TestInput[string]{
				haystack: []string{"1", "2", "3", "4", "5"},
				needle:   "7",
			},
			-1,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result := Linear[string](&test.input.haystack, test.input.needle)
			assert.Equal(t, test.result, result)
		})
	}
}
