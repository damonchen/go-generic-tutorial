package main

import (
	"fmt"
	"strconv"
)

type Setter interface {
	Set(string)
}

// func FromString[T Setter](s []string) []T {
// 	// result := make([]T, len(s))

// 	var result []T
// 	for _, v := range s {
// 		t := T()
// 		t := t.New()
// 		t.Set(v)
// 		append(result, t)
// 	}

// 	return result
// }

type Settable int

func (p *Settable) Set(s string) {
	i, _ := strconv.Atoi(s)
	// 由于T是指针，所以指针值是nil，此处会panic
	*p = Settable(i)
}

// func (p *Settable) New() *Settable {
// 	return &Settable{}
// }

type Setter2[B any] interface {
	Set(string)
	*B // non-interface type constraint element
}

func FromStrings2[T any, PT Setter2[T]](s []string) []T {
	result := make([]T, len(s))
	for i, v := range s {
		// The type of &result[i] is *T which is in the type set
		// of Setter2, so we can convert it to PT.
		p := PT(&result[i])
		// PT has a Set method.
		p.Set(v)
	}
	return result
}

func main() {
	// nums := FromString[*Settable]([]string{"1", "2"})
	nums := FromStrings2[Settable]([]string{"1", "2"})
	fmt.Println(nums)
}
