package main

import (
    "sort"
    "fmt"
)

// sliceFn is an internal type that implements sort.Interface.
// The Less method calls the cmp field.
type sliceFn[T any] struct {
	s   []T
	cmp func(T, T) bool
}

func (s sliceFn[T]) Len() int           { return len(s.s) }
func (s sliceFn[T]) Less(i, j int) bool { return s.cmp(s.s[i], s.s[j]) }
func (s sliceFn[T]) Swap(i, j int)      { s.s[i], s.s[j] = s.s[j], s.s[i] }

// SliceFn sorts the slice s according to the function cmp.
func SliceFn[T any](s []T, cmp func(T, T) bool) {
	sort.Sort(sliceFn[T]{s, cmp})
}

type Person struct {
    Name string
}

func (p *Person) String() string {
    return p.Name
}


func main() {
    var s []*Person
    s = append(s, &Person{
        Name: "zoo",
    })
    s = append(s, &Person{
        Name: "damon",
    })
    s = append(s, &Person{
        Name: "operation",
    })

    SliceFn(s, func(p1, p2 *Person) bool { return p1.Name < p2.Name })

    fmt.Println(s)
}
