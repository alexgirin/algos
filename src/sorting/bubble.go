package sorting

import "cmp"

func BubbleSort[T cmp.Ordered](array *[]T) {
	pointer := *array
	for i := 0; i < len(pointer); i++ {
		for j := 0; j < len(pointer)-1-i; j++ {
			if pointer[j] > pointer[j+1] {
				pointer[j], pointer[j+1] = pointer[j+1], pointer[j]
			}
		}
	}
}
