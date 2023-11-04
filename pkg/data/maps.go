package data

// IndexSlice creates a map from the slice sl keyed by the string returned by indexer for each
// element in sl. Duplicated elements are removed.
func IndexSlice[S ~[]E, E any](sl S, indexer func(E) string) map[string]E {
	m := make(map[string]E, len(sl))

	for _, el := range sl {
		key := indexer(el)
		m[key] = el
	}

	return m
}

// MapToSlice converts a map to a slice. The order of the slice elements
// is unspecified.
func MapToSlice[S ~map[K]V, K comparable, V any](m S) []V {
	res := make([]V, 0, len(m))

	for _, val := range m {
		res = append(res, val)
	}

	return res
}
