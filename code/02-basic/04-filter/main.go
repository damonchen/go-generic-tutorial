package main

import (
	"fmt"
)

func Filter[T any](s []T, f func(T) bool) []T {
  var r []T
  for _, v := range s {
    if f(v) {
      r = append(r, v)
    }
  }
  return r
}



func main() {
	s := []int{1, 2, 3, 4, 5}
	v := Filter(s, func(a int) bool {
		return a != 3
	})
	fmt.Println(v)
}