package gocol

func GroupBy[T any, K comparable](original []T, groupKey func(T) K) map[K][]T {
	if original == nil { return nil }
	res := make(map[K][]T)
	for _, t := range original {
		res[groupKey(t)] = append(res[groupKey(t)], t)
	}
	return res
}

func GroupMap[T any, V any, K comparable](original []T, groupKey func(T) K, mapFunc func(T) V) map[K][]V {
	if original == nil { return nil }
	res := make(map[K][]V)
	for _, t := range original {
		res[groupKey(t)] = append(res[groupKey(t)], mapFunc(t))
	}
	return res
}

