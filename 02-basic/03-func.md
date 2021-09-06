# 进阶

## 函数接口

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

## 操作符

对于一些比较操作符、算术符号以及逻辑操作符，一般只有预定义的基础类型支持，上面的函数接口就无法满足了。

如
```go
func Less[T any](a, b T) bool {
  return a < b
}
```

由于`T`是`any`类型，并没有`<`的支持，所以上述代码编译会出问题


### 类型集合

每一个类型，有一个关联的类型集合，对于非接口类型的T，其集合为 {T}，包含自己的集合。

对于任意方法签名`M1`和`M2`,`interface{M1; M2}`表示的所有满足`M1`和`M2`的类型，表示的是`M1`和`M2`集合的交集。

比方说下面的示例：

```go

type File interface {
  Read(p []byte) (n int, err error)
  Write(p []byte) (n int, err error)
  Close() error
}

type ReaderFiler interface {
  File;
  io.ReaderCloser()
}

func read[T ReaderFiler](f T) (n int, err error) {

}

```


### Approximation constraint element

如果我们已经定义了`type MyString string`，如果在接口定义中，我们使用了如下的代码：

```go
type AnyString interface {
  string
}
```

那么`MyString`尽管理论上是满足上面的约束的，但由于实现的缘故，不能够匹配，所以，`go`引入了`~`操作符，以此来实现：

- The type set of `~T` is the set of all types whose underlying type is `T`.

这里面包含了两个意思：

1. Type literal, 如[]byte, struct {f int}
2. 大部分预定义的类型，如int或string,但不包括error

`~T`在`T`为类型参数或者`T`是接口类型的时候是会报错的


### Union constraint element

通过`|`来表示或的方式:

```go
type PredeclaredSignedInteger interface {
	int | int8 | int16 | int32 | int64
}
```

表示的是类型是{int, int8, int16, int32, int64}的集合。

当然接口中的类型，也可以是类似 `~int` 这样形式的


### 类型支持的操作

大小比较操作，则仅对数值和字符串有效

相等比较，则在上面的基础上，还包含：结构体、数组以及接口类型

### 相互引用参数

在图的数据结构中，边(Edge)会引用点(Node)，点(Node)会引用边(Edge)：

```go
package graph

// NodeConstraint is the type constraint for graph nodes:
// they must have an Edges method that returns the Edge's
// that connect to this Node.
type NodeConstraint[Edge any] interface {
	Edges() []Edge
}

// EdgeConstraint is the type constraint for graph edges:
// they must have a Nodes method that returns the two Nodes
// that this edge connects.
type EdgeConstraint[Node any] interface {
	Nodes() (from, to Node)
}

// Graph is a graph composed of nodes and edges.
type Graph[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]] struct { ... }

// New returns a new graph given a list of nodes.
func New[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]] (nodes []Node) *Graph[Node, Edge] {
	...
}

// ShortestPath returns the shortest path between two nodes,
// as a list of edges.
func (g *Graph[Node, Edge]) ShortestPath(from, to Node) []Edge { ... }
```


### 类型推理

为了简便，不直接写出所有的参数类型,可以通过函数参数类型来推理从非类型参数类型来推断类型参数。我们可以使用约束类型推理从已知类型参数推断未知类型参数。

```
func Map[F, T any](s []F, f func(F) T) []T {
  ...
}

var s []int
var r []int64
f := func(i int) int64 { return int64(i) }
r = Map(s, f)

```

如果系统不能够推断出类型，那么系统就会报错。

通过推断的方式，不增加语法，但代码更容易读。

`[]map[int]bool`可以匹配的类型有：

```go
[]map[int]bool
T1 (T1 matches []map[int]bool)
[]T1 (T1 matches map[int]bool)
[]map[T1]T2 (T1 matches int, T2 matches bool)
```

类型推导中：

```
// NewPair returns a pair of values of the same type.
func NewPair[F any](f1, f2 F) *Pair[F] { ... }
```

里面的结构要求F参数是相同的类型，所以这样的代码：

`NewPair(1, 2.5)`

会出错


When constraint type inference is possible, type inference proceeds as followed(类型推到):

- Build the mapping using known type arguments. (类型参数匹配)
- Apply constraint type inference. （通过约束来进行类型推导）
- Apply function type inference using typed arguments.（通过类型参数来进行函数类型推导）
- Apply constraint type inference again.（通过约束来进行类型推导）
- Apply function type inference using the default types of any remainint untyped arguments.（对未确定的类型参数使用默认的类型进行函数类型推导）
- Apply constraint type inference again. （通过约束来进行类型推导）

