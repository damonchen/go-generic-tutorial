package main

import (
	"fmt"
)

func p[T any](a T) string {
  return fmt.Sprintf("%+v", a)
}

type Person struct {
	Name string
}

func main() {
	fmt.Println(p("hello"))
	fmt.Println(p(1))
	fmt.Println(p(uint32(2)))
	fmt.Println(p(uint64(3)))
	fmt.Println(p(Person{Name:"world"}))
}