package gocol

func Filter[T any](original []T, predicate func(T) bool) []T {
	if original == nil {
		return nil
	}
	var res []T
	for _, v := range original {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}

func FilterNot[T any](original []T, predicate func(T) bool) []T {
	if original == nil {
		return nil
	}
	var res []T
	for _, v := range original {
		if !predicate(v) {
			res = append(res, v)
		}
	}
	return res
}

func Collect[F any, T any](original []F, predicate func(F) bool, mapFunc func(F) T) []T {
	if original == nil {
		return nil
	}
	var res []T
	for _, v := range original {
		if predicate(v) {
			res = append(res, mapFunc(v))
		}
	}
	return res
}
