# 第一个泛型程序


## hello world

生成下面的文件`hello.go`

```go
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

```

执行：
```bash
$ gotip run hello.go
3
hello world
```

`go`泛型的核心在于类型约束，即下面描述的类型

```go
type AddType interface{
	~int |
	~string
}
```

`~int`是一个新的语法，表示的是所有底层类型为`int`的类型的集合，一个近似的元素的表述。
