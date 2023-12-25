package search

func Linear[T comparable](haystack *[]T, needle T) int {
	for index, val := range *haystack {
		if val == needle {
			return index
		}
	}

	return -1
}
