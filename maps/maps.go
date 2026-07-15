package maps

// InvertMap swaps keys and values.
// The value type (V) must be comparable to be used as a map key.
func InvertMap[K comparable, V comparable](original map[K]V) map[V]K {
	inverted := make(map[V]K, len(original))
	for k, v := range original {
		inverted[v] = k
	}
	return inverted
}
