package gocol

func Map[F any, T any](original []F, mapFunc func(F) T) []T {
	if original == nil {
		return nil
	}
	length := len(original)
	res := make([]T, length)
	for i, f := range original {
		res[i] = mapFunc(f)
	}
	return res
}

func MapInPlace[T any](original []T, mapFunc func(T) T) {
	for i, v := range original {
		original[i] = mapFunc(v)
	}
}

func FlatMap[F any, T any](original []F, flatMapFunc func(F) []T) []T {
	res := []T{}
	for _, f := range original {
		res = append(res, flatMapFunc(f)...)
	}
	return res
}

func DistinctUnordered[T comparable](original []T) []T {
	unique := make(map[T]struct{})
	for i := range original {
		unique[original[i]] = struct{}{}
	}
	result := []T{}
	for k := range unique {
		result = append(result, k)
	}
	return result
}

func Forall[T any](original[]T, predicate func(T) bool) bool {
    if original == nil { return false }
    predicateHolds := true
    for _, value := range original {
        if !predicate(value) {
            predicateHolds = false
            break
        }
    }
    return predicateHolds
}
