package gocol

import (
    "testing"
)

func TestFirst(t *testing.T) {
    slice := []int{1, 7, 3, 6, 1, 9}
    expected := slice[0]
    actual := Reduce(slice, First[int]())
    if expected != actual {
        t.Errorf("first argument did not match the expected value. expected %d got %d", expected, actual)
    }
}

func TestLast(t *testing.T) {
    slice := []int{1, 6, 2, 7, 2, 9, 0, 1, 4, 7}
    expected := slice[len(slice) - 1]
    actual := Reduce(slice, Last[int]())
    if expected != actual {
        t.Errorf("last argument did not match the expected value. expected %d got %d", expected, actual)
    }
}

func TestGroupBy(t *testing.T) {
    slice := []int{1, 6, 2, 4, 6, 8, 3, 2, 7, 9, 2}
    even := func(i int) bool { return i % 2 == 0 }
    result := GroupBy(slice, even)

    odd := func(i int) bool { return i % 2 != 0 }
    for isEven, values := range result {
        if !((isEven && Forall(values, even)) || (!isEven && Forall(values, odd))) {
            t.Error("GroupBy did not work as expected")
        }
    }
}
