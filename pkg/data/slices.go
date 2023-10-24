package data

import "slices"

// SliceOverlaps returns true if at least one element is present in both
// slices a and b.
func SliceOverlaps[S ~[]E, E comparable](a S, b S) bool {
	for _, el := range a {
		if slices.Contains(b, el) {
			return true
		}
	}

	return false
}

// SliceOverlapsFunc returns true if at least one element is present in both
// slices a and b.
func SliceOverlapsFunc[A ~[]E1, E1 comparable, B ~[]E2, E2 comparable](a A, b B, comp func(el E2) E1) bool {
	for _, el2 := range b {
		convEl1 := comp(el2)
		if slices.Contains(a, convEl1) {
			return true
		}
	}

	return false
}
