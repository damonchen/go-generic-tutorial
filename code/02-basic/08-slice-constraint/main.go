package main

import "fmt"

type SliceConstraint[T any] interface {
	[]T
}

func Map[S SliceConstraint[E], E any](s S, f func(E) E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

type MySlice []int


func DoubleMySlice(s MySlice) MySlice {
	v := Map[MySlice, int](s, func(e int) int { return 2 * e })
	return v
}

func main() {
	s := MySlice{1, 2, 3}
	d := DoubleMySlice(s)
	fmt.Println(d)
}
