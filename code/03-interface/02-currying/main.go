package main

import (
    "fmt"
)


type Integer interface {
    ~int|~uint8|~uint64
}

type UFunc func[U any](U) U

func addAndMulti[T Integer, U Integer](a T) UFunc {
    b := 10 + a
    return multi(u U) U {
        return b * u
    }
}


func main() {
    f := addAndMulti(10)
    fmt.Println(f(3))
}
