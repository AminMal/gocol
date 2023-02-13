package gocol

import (
	"fmt"
)

func Fold[T any, V any](original []T, zero V, reduce func(V, T) V) V {
	if original == nil {
		return zero
	}
	res := zero
	for i := range original {
		res = reduce(res, original[i])
	}
	return res
}

func FoldRight[T any, V any](original []T, zero V, reduce func(V, T) V) V {
	if original == nil {
		return zero
	}
	res := zero
	for i := len(original); i > 0; i-- {
		res = reduce(res, original[i-1])
	}
	return res
}

func emptySliceReduce() { panic(fmt.Errorf("reduce on empty slice")) }

func Reduce[T any](original []T, reduce func(T, T) T) T {
	if original == nil {
		emptySliceReduce()
	}
	return Fold(original[1:], original[0], reduce)
}

func ReduceRight[T any](original []T, reduce func(T, T) T) T {
	if original == nil {
		emptySliceReduce()
	}
	return FoldRight(original[:len(original)-1], original[len(original)-1], reduce)
}
