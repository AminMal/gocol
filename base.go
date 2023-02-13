package gocol

func Identity[T any](t T) T { return t }

func NoOp[T any](t T) T { return t }

type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}
