# 基础

读本文之前，要求读者对`go`这门语言有一定的基础


## 新的内容

泛型相比与之前的`go`，引入了几个特殊的内容。

**首先**是泛型的表示方式，和`Java`、`C++`、`C#`、`Swift`等语言的泛型表述方式不一样，`go`的泛型，没有采用`<>`来表述，采用的是`[]`来表述，如下面代码中的`[T AddType]`：

```go
func add[T AddType](a, b T) T {
  return a + b
}
```

其中`AddType`为泛型`T`的类型约束

之所以没有采用`<>`的原因是因为`go`的函数支持多返回值，`<>`存在二义性，如下:

```go
a, b = w < x, y > (z)
```
上面的代码中，到底是`w<x`和`y>(z)`的表述还是泛型的表述？


**其次**和`Java`、`C++`、`C#`、`Swift`等语言不一样的地方是，`go`采用的是类型参数(Type Parameter)的方式来实现的，也就是上面示例中的`AddType`:

```go
type AddType interface{
	~int |
	~string
}
```

这儿，go引入了两个操作符：`~`和`|`：

- `~T` 表示的是底层类型是T，是一个约等于的表述方式
- `|` 表示的是一个或的意思，表示类型可以是其中的一个

前面的`AddType`是一个类型约束的定义，其中类型约束定义要求必须是`interface`类型。`AddType`的类型约束可以这么描述：满足`int`或者`string`类型的那些类型。

为了能够对更一般情形的类型约束进行描述，`go`引入了`any`作为任意类型的约束(实际上等价于`interface{}`，类似rune和byte)，如下面的代码中的`T`用`any`来约束，表示`T`可以为任意类型:

```go
func p[T any](a T) string {
  return fmt.Sprintf("%v", a)
}
```


## 小结

- 通过`interface`来定义类型约束(如`AddType`)，新增了`any`的标识符，用来描述所有类型
- 通过`[T AddType]`来进行`T`的约束描述


## 入门示例

类似`JavasSript`或者`Python`中的`map`、`filter`和`reduce`实现

### map

```go
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
  r := make([]T2, len(s))
  for i, v := range s {
    r[i] = f(v)
  }
  return r
}
```

### filter

```go
func Filter[T any](s []T, f func(T) bool) []T {
  var r []T
  for _, v := range s {
    if f(v) {
      r = append(r, v)
    }
  }
  return r
}
```


### reduce

```go
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
  r := initializer
  for _, v := range s {
    r = f(r, v)
  }
  return r
}
```

