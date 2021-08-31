# 进阶

在前面的章节中，我们描述了，泛型的约束是通过`interface`来定义的，那么自然而然的，
`go`中接口的描述，还可以涉及函数，所以下面的`Stringer`也可以作为类型约束来使用

```go
type Stringer interface {
  String() string
}
```

将数组字符串化

```go
func Stringify[T Stringer](s []T) []string {
  r := make([]string, len(s))
  for i, v := range s {
    r[i] = v.String()
  }
  return r
}
```

这儿就会要求调用`Stringify`函数的参数，必须是满足`Stringer`接口类型的数组