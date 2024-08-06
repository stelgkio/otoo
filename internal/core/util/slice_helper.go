package util

import "sort"

// RemoveElement removes an element from a slice at a given index
func RemoveElement[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

// RemoveElements removes multiple elements from a slice at given indices
func RemoveElements[T any](s []T, indices []int) []T {
	sort.Sort(sort.Reverse(sort.IntSlice(indices)))
	for _, index := range indices {
		if index >= 0 && index < len(s) {
			s = append(s[:index], s[index+1:]...)
		}
	}
	return s
}

// FindIndex returns the index of the first occurrence of the element in the slice
func FindIndex[T comparable](s []T, element T) int {
	for i, v := range s {
		if v == element {
			return i
		}
	}
	return -1
}
