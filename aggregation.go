package gocol

func Add[T Numeric]() func(T, T) T {
	return func(a, b T) T { return a + b }
}

func First[T any]() func(T, T) T {
	return func(head, _ T) T { return head }
}

func Last[T any]() func(T, T) T {
	return func(_, last T) T { return last }
}

func GroupBy[T any, K comparable](original []T, groupKey func(T) K) map[K][]T {
	if original == nil {
		return nil
	}
	res := make(map[K][]T)
	for _, t := range original {
		res[groupKey(t)] = append(res[groupKey(t)], t)
	}
	return res
}

func GroupMap[T any, V any, K comparable](original []T, groupKey func(T) K, mapFunc func(T) V) map[K][]V {
	if original == nil {
		return nil
	}
	res := make(map[K][]V)
	for _, t := range original {
		res[groupKey(t)] = append(res[groupKey(t)], mapFunc(t))
	}
	return res
}

func GroupMapReduce[T any, V any, K comparable](original []T, groupKey func(T) K, mapFunc func(T) V, reduce func (V, V) V) map[K]V {
	if original == nil { return nil }
	res := make(map[K]V)

	for i := range original {
		t := original[i]
		res[groupKey(t)] = reduce(res[groupKey(t)], mapFunc(t))
	}
	return res
}
