package util

import (
	"reflect"
	"testing"
)

func TestFindIndex(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"a", "b", "c", "d", "e"}

	// Test cases for int slice
	if got := FindIndex(intSlice, 3); got != 2 {
		t.Errorf("FindIndex(intSlice, 3) = %d; want 2", got)
	}
	if got := FindIndex(intSlice, 10); got != -1 {
		t.Errorf("FindIndex(intSlice, 10) = %d; want -1", got)
	}

	// Test cases for string slice
	if got := FindIndex(stringSlice, "c"); got != 2 {
		t.Errorf("FindIndex(stringSlice, 'c') = %d; want 2", got)
	}
	if got := FindIndex(stringSlice, "z"); got != -1 {
		t.Errorf("FindIndex(stringSlice, 'z') = %d; want -1", got)
	}
}

func TestRemoveElement(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	wantIntSlice := []int{1, 2, 4, 5}

	// Test removing an element from int slice
	if got := RemoveElement(intSlice, 2); !reflect.DeepEqual(got, wantIntSlice) {
		t.Errorf("RemoveElement(intSlice, 2) = %v; want %v", got, wantIntSlice)
	}

	// Test removing an element with an out of bounds index
	if got := RemoveElement(intSlice, 10); !reflect.DeepEqual(got, intSlice) {
		t.Errorf("RemoveElement(intSlice, 10) = %v; want %v", got, intSlice)
	}
}

func TestRemoveElements(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	indicesToRemove := []int{1, 3}
	wantIntSlice := []int{1, 3, 5}

	// Test removing multiple elements from int slice
	if got := RemoveElements(intSlice, indicesToRemove); !reflect.DeepEqual(got, wantIntSlice) {
		t.Errorf("RemoveElements(intSlice, indicesToRemove) = %v; want %v", got, wantIntSlice)
	}

	// Test removing elements with an out of bounds index
	intSliceOutOfBounds := []int{1, 2, 3, 4, 5}
	indicesToRemoveOutOfBounds := []int{1, 10}
	wantIntSliceOutOfBounds := []int{1, 3, 4, 5}
	if got := RemoveElements(intSliceOutOfBounds, indicesToRemoveOutOfBounds); !reflect.DeepEqual(got, wantIntSliceOutOfBounds) {
		t.Errorf("RemoveElements(intSliceOutOfBounds, indicesToRemoveOutOfBounds) = %v; want %v", got, wantIntSliceOutOfBounds)
	}
}
