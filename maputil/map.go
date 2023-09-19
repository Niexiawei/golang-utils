package maputil

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))

	var i int
	for _, v := range m {
		values[i] = v
		i++
	}

	return values
}
