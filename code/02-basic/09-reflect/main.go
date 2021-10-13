package main

import (
    "fmt"
    "reflect"
)

type stack [T any] []T


func (s *stack[T]) push(e T) {
    *s = append(*s, e)
}


type Stack[T any] struct {
    values []T
}

func (s *Stack[T]) push(e T) {
    s.values = append(s.values, e)
}

type Integer interface {
    ~int|~uint|~uint8|~uint16|~uint32
}

func min[T Integer] (a, b T) T {
    if a < b {
        return a
    }
    return b
}

func main() {
    // s := stack[int]{}
    // s.push(10)
    // fmt.Println(s)

    // fmt.Println(reflect.TypeOf(s))

    s := Stack[int]{}
    s.push(10)
    s.push(20)
    s.push(30)
    fmt.Println(reflect.TypeOf(s))

    fmt.Println(min(10, 20))


    for _, t := range s.values {
        fmt.Println(t)
    }
}
