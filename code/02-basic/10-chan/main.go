package main

import (
	"fmt"
)

type Stack[T any] struct {
	v chan T
}

func (s Stack[T]) Len() int { return len(s.v)}

func (s Stack[T]) push(v T) { s.v <- v}

func (s Stack[T]) pop() T {
	v := <- s.v
	return v
}

func main() {
	s := Stack{
		v: make(chan int),
	}

	go func() {
		s.push(10)
	}()

	c := s.pop()
	fmt.Println(c)
}
