package slice

func UniqueValues[T any, R comparable](items []T, transform func(T) R) []R {
	m := map[R]struct{}{}
	for _, item := range items {
		m[transform(item)] = struct{}{}
	}

	uniqueValues := make([]R, 0, len(m))
	for v := range m {
		uniqueValues = append(uniqueValues, v)
	}

	return uniqueValues
}
