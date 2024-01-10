package sorting

import (
	"cmp"
)

func merge[T cmp.Ordered](left []T, right []T) []T {
	var i, j int
	lenLeft, lenRight := len(left), len(right)
	lenResult := lenLeft + lenRight
	result := make([]T, lenResult)

	for k := 0; k < lenResult; k++ {
		if i > lenLeft-1 && j <= lenRight-1 {
			result[k] = right[j]
			j++
		} else if j > lenRight-1 && i <= lenLeft-1 {
			result[k] = left[i]
			i++
		} else if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}

	return result
}

func MergeSort[T cmp.Ordered](array *[]T) []T {
	length := len(*array)
	if length <= 1 {
		return *array
	}

	mid := length / 2

	left := (*array)[:mid]
	right := (*array)[mid:]

	return merge(MergeSort[T](&left), MergeSort[T](&right))
}
