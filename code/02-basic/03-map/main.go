package main

import (
	"fmt"
)

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
  r := make([]T2, len(s))
  for i, v := range s {
    r[i] = f(v)
  }
  return r
}


func main() {
	s := []int{1, 2, 3, 4, 5}
	v := Map(s, func(a int) int {
		return a * 2
	})
	fmt.Println(v)
}