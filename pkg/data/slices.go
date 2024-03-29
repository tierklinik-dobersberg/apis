package data

import "slices"

// SliceOverlaps returns true if at least one element is present in both
// slices a and b.
//
// Deprecated: use ElemInBothSlices
func SliceOverlaps[S ~[]E, E comparable](a S, b S) bool {
	return ElemInBothSlices(a, b)
}

// ElemInBothSlices returns true if at least one element is present in both
// slices a and b.
func ElemInBothSlices[S ~[]E, E comparable](a S, b S) bool {
	for _, el := range a {
		if slices.Contains(b, el) {
			return true
		}
	}

	return false
}

// SliceOverlapsFunc returns true if at least one element is present in both
// slices a and b by using a function conv to convert elements from b to the same type
// of elements in a.
//
// Deprecated: Use ElemInBothSlicesFunc
func SliceOverlapsFunc[A ~[]E1, E1 comparable, B ~[]E2, E2 comparable](a A, b B, conv func(el E2) E1) bool {
	return ElemInBothSlicesFunc(a, b, conv)
}

// ElemInBothSlicesFunc returns true if at least one element is present in both
// slices a and b by using a function conv to convert elements from b to the same type
// of elements in a.
func ElemInBothSlicesFunc[A ~[]E1, E1 comparable, B ~[]E2, E2 comparable](a A, b B, conv func(el E2) E1) bool {
	for _, el2 := range b {
		convEl1 := conv(el2)
		if slices.Contains(a, convEl1) {
			return true
		}
	}

	return false
}
