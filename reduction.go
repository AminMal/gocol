package gocol

import (
	"fmt"
)

// todo, WIP
// func Fold[T any](original []T

func emptySliceReduce() { panic(fmt.Errorf("reduce on empty slice")) }

func Reduce[T any](original []T, reduce func (T, T) T) T {
	if original == nil { emptySliceReduce() }
	var t T
	for _, v := range original {
		t = reduce(t, v)
	}
	return t
}

func ReduceRight[T any](original []T, reduce func(T, T) T) T {
	if original == nil { emptySliceReduce() }
	var t T
	for _, v := range original {
		t = reduce(v, t)
	}
	return t
}

