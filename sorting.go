package gocol

import (
	"sort"
)

func cpy[T any](src []T) []T {
	result := []T{}
	for _, t := range src {
		result = append(result, t)
	}
	return result
}

func Sorted[T any](ts []T, o Ordering[T]) []T {
	clone := cpy(ts)
	sort.Slice(clone, func(i, j int) bool {
		return o.Compare(clone[i], clone[j]) < 0
	})
	return clone
}

func SortInPlace[T any](ts []T, o Ordering[T]) {
	sort.Slice(ts, func(i, j int) bool {
		return o.Compare(ts[i], ts[j]) < 0
	})
}

func OrderBy[T any, K any](ts []T, convert func(T) K, kOrdering Ordering[K]) []T {
	tOrd := OrderingBasedOn(kOrdering, convert)
	return Sorted(ts, tOrd)
}

func OrderByInPlace[T any, K any](ts []T, convert func(T) K, kOrdering Ordering[K]) {
	tOrd := OrderingBasedOn(kOrdering, convert)
	SortInPlace(ts, tOrd)
}
