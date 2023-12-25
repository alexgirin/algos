package search

import "cmp"

func Linear[T cmp.Ordered](haystack *[]T, needle T) int {
	for index, val := range *haystack {
		if val == needle {
			return index
		}
	}

	return -1
}
