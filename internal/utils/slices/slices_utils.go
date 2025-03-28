package slices

func FilterOut[T comparable](source []T, remove []T) []T {
	// Crear un set para rápida búsqueda
	removeSet := make(map[T]struct{})
	for _, v := range remove {
		removeSet[v] = struct{}{}
	}

	var result []T
	for _, v := range source {
		if _, found := removeSet[v]; !found {
			result = append(result, v)
		}
	}
	return result
}

func FilterByField[T any](items []T, exclude []string, selector func(T) string) []T {
	excludeSet := make(map[string]struct{}, len(exclude))
	for _, id := range exclude {
		excludeSet[id] = struct{}{}
	}

	var result []T
	for _, item := range items {
		if _, found := excludeSet[selector(item)]; !found {
			result = append(result, item)
		}
	}
	return result
}

func MapToProp[T any](items []T, selector func(T) string) []string {
	var p []string
	for _, item := range items {
		p = append(p, selector(item))
	}
	return p
}
