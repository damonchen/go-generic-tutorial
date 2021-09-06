package main

import (
	"fmt"
)

type Integer interface {
	~int
}

func add(a, b Integer) Integer {
	return a + b
}

func main() {
	// c := []int{1,2}
	// d := []int{3,4}

	// e := add(c, d)
    c := int(3)
    d := int(4)
    e := add(c, d)
	fmt.Println(e)
}
