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
