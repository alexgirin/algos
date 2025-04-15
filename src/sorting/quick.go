package sorting

import "cmp"

func partition[T cmp.Ordered](array []T, low, high int) int {
	index := low - 1
	pivotElement := array[high]
	for i := low; i < high; i++ {
		if array[i] <= pivotElement {
			index += 1
			array[index], array[i] = array[i], array[index]
		}
	}
	array[index+1], array[high] = array[high], array[index+1]
	return index + 1
}

func quickSortRange[T cmp.Ordered](array []T, low, high int) {
	if len(array) <= 1 {
		return
	}

	if low < high {
		mid := partition(array, low, high)
		quickSortRange(array, low, mid-1)
		quickSortRange(array, mid+1, high)
	}
}

func QuickSort[T cmp.Ordered](array *[]T) []T {
	quickSortRange(*array, 0, len(*array)-1)
	return *array
}
