package search

import "cmp"

func Binary[T cmp.Ordered](haystack *[]T, needle T) int {
	low, high := 0, len(*haystack)-1

	for low <= high {
		mid := low + (high-low)/2

		if needle == (*haystack)[mid] {
			return mid
		}

		if needle > (*haystack)[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return -1
}
