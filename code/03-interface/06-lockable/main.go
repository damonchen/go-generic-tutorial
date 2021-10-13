package main

import (
	"fmt"
	"sync"
)

type Lockable[T any] struct {
	T
	mu sync.Mutex // guards
}

func (l *Lockable[T]) Get() T {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.T
}

func (l *Lockable[T]) Set(v T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.T = v
}


// type NamedInt[Name fmt.Stringer] struct {
// 	Name
// 	val int
// }

// // Name returns the name of a NamedInt.
// func (ni NamedInt[Name]) Name() string {
// 	// The String method is promoted from the embedded Name.
// 	return ni.String()
// }

// type P interface {
// 	Name() string
// }

type Person struct {
	Name string
}

// func (p Person) Name() string {
// 	return "Person"
// }



func Switch2[T any](v interface{}) int {
	switch v.(type) {
	case T:
		return 0
	case string:
		return 1
	default:
		return 2
	}
}


func main() {
    // l := &Lockable[int]{
       
    // }

    // p := Person{}
 //    p := int(10)
	// l.Set(p)

	// fmt.Println(l.Get())


	b := int(100)
	a := Switch2(b)
	fmt.Println(a)


}
