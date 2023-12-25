package search

func Linear[T comparable](haystack *[]T, needle T) bool {
	for _, val := range *haystack {
		if val == needle {
			return true
		}
	}

	return false
}
