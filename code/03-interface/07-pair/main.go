package main

import (
    "fmt"
)


type Pair[T any] struct {
    f1, f2 T
}

func main() {
    a := int(1)
    b := int(2)
    v := Pair[int]{a, b}

    fmt.Println(v)

}
