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
