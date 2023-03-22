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

func Find[F any](original []F, predicate func(F) bool) (F, bool) {
	var result F
	found := false
	for _, v := range original {
		if predicate(v) {
			result = v
			found = true
			break
		}
	}
	return result, found
}

func Exists[F any](original []F, predicate func(F) bool) bool {
	found := false
	for _, v := range original {
		if predicate(v) {
			found = true
			break
		}
	}
	return found
}
