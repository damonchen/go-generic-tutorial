package main

import (
    "fmt"
)

type Optional[T any] struct {
	p *T
}

func NewOptional[T any](v *T) Optional[T] {
    return Optional[T]{
        p: v,
    }
}

func (o Optional[T]) Val() T {
	if o.p != nil {
		return *o.p
	}
	var zero T
	return zero
}


func IntToT[T any](c int) T {
    var t T
    t = (T)(c)
    return t
}


func GetOptional[T any]() Optional[T] {
    v := IntToT[T](10)
    o := NewOptional[T](&v)
    return o
}

func main() {
    o := GetOptional[int]()
    fmt.Println(o.Val())

    c := "vvvvvv"
    o2 := Optional[string]{
        p: &c,
    }
    fmt.Println(o2.Val())
}

