package gocol

import (
	"time"
)

type Ordering[T any] interface {
	Compare(one, another T) int
}

type genericOrdering[T any] struct {
	compareFunc func(a, b T) int
}

func (g genericOrdering[T]) Compare(one, another T) int { return g.compareFunc(one, another) }

// ------- Ordering composition -------
type using[T any, V any] struct {
	original Ordering[T]
	convertFunc func(V)T
}

func (u using[T, V]) Compare(a, b V) int {
	return u.original.Compare(u.convertFunc(a), u.convertFunc(b))
}

func OrderingBasedOn[T any, V any](original Ordering[T], convert func(V)T) Ordering[V] {
	return using[T, V]{original, convert}
}

func NumericAsc[T Numeric]() Ordering[T] {
	compare := func(a, b T) int {
		if a < b { return -1 }
		if a == b { return 0 }
		return 1
	}
	return genericOrdering[T]{compare}
}

func NumericDesc[T Numeric]() Ordering[T] {
	compare := func(a, b T) int {
		if a < b { return 1 }
		if a == b { return 0 }
		return -1
	}
	return genericOrdering[T]{compare}
}

// ------- Numeric Instances (for ease of use) -------

var IntAsc Ordering[int] = NumericAsc[int]()

var IntDesc Ordering[int] = NumericDesc[int]()

var UintAsc Ordering[uint] = NumericAsc[uint]()

var UintDesc Ordering[uint] = NumericDesc[uint]()

var Float32Asc Ordering[float32] = NumericAsc[float32]()

var Float32Desc Ordering[float32] = NumericDesc[float32]()

var Float64Asc Ordering[float64] = NumericAsc[float64]()

var Float64Desc Ordering[float64] = NumericDesc[float64]()

// ------- Date&Time Orderings -------
var TimeAsc Ordering[time.Time] = genericOrdering[time.Time]{
	compareFunc: func(a, b time.Time) int {
		if a.Before(b) { return -1 }
		if a.After(b) { return 1 }
		return 0
	},
}

var TimeDesc Ordering[time.Time] = genericOrdering[time.Time]{
	compareFunc: func(a, b time.Time) int {
		if a.After(b) { return -1 }
		if a.Before(b) { return 1 }
		return 0
	},
}
