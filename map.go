package gocol

func Map[F any, T any](original []F, mapFunc func(F)T) []T {
	if original == nil { return nil }
	length := len(original)
	res := make([]T, length, length)
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

func Bind[F any, T any](original []F, bind func(F) []T) []T {
	return FlatMap[F, T](original, bind)
}

