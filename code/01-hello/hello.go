package main

import (
  "fmt"
)

type AddType interface{
	~int |
	~string
}

func add[T AddType](a, b T) T {
  return a + b
}

func main() {
  fmt.Println(add(1,2))
  fmt.Println(add("hello", " world"))
}
