package main

import (
	"fmt"
)

type Stringer interface {
  String() string
}

func Stringify[T Stringer](s []T) []string {
  r := make([]string, len(s))
  for i, v := range s {
    r[i] = v.String()
  }
  return r
}

type Person struct {
	Name string
}

func (p Person) String() string {
	return "hello " + p.Name
}

type Int int

func (i Int) String() string {
	return fmt.Sprintf("hello %d", i)
}

func main() {
	v := []Stringer{Int(42), Person{Name: "world"}}
	s := Stringify(v)
	for _, e := range s {
		fmt.Println(e)
	}
}