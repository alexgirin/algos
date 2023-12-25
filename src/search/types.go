package search

type TestInput[T comparable] struct {
	haystack []T
	needle   T
}
