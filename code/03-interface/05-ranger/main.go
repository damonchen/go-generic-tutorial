package main


import (
    "fmt"
    "runtime"
)


type Sender[T any] struct {
    values chan<- T
    done <- chan bool
}

func (s *Sender[T]) Send(v T) bool {
    select {
    case s.values <- v:
        return true
    case <- s.done:
        return false
    }
}

func (s *Sender[T]) Close() {
    close(s.values)
}

type Receiver[T any] struct {
    values <- chan T
    done chan <- bool
}


func (r *Receiver[T]) Next() (T, bool) {
    v, ok := <- r.values
    return v, ok
}

func (r *Receiver[T]) finalize() {
    close(r.done)
}

func Ranger[T any]() (*Sender[T], *Receiver[T]) {
    c := make(chan T)
    d := make(chan bool)

    s := &Sender[T]{values: c, done: d}

    r := &Receiver[T]{values: c, done: d}

    runtime.SetFinalizer(r, func(r *Receiver[T]) { r.finalize() })
    // runtime.KeepAlive(r)
    return s, r
}

func main() {
    s, r := Ranger[int]()

    done := make(chan bool)
    go func() {
        fmt.Println("send 1", s.Send(1))
        fmt.Println("send 2", s.Send(2))
        fmt.Println("send 3", s.Send(3))
    }()


    fmt.Println(runtime.Version())


    go func() {
        b, ok := r.Next()
        fmt.Println("receive", b, ok)
        b, ok = r.Next()
        fmt.Println("receive", b, ok)

        done <- true
    }()

    <- done

    fmt.Println("done")
}
