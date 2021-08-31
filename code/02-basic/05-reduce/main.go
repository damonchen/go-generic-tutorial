package main

import (
	"fmt"
)

func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
  r := initializer
  for _, v := range s {
    r = f(r, v)
  }
  return r
}


func main() {
	s := []int{1, 2, 3, 4, 5}
	v := Reduce(s, 0, func(p, c int) int {
		return p + c
	})
	fmt.Println(v)
}