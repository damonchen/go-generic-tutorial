# faq

Q:

  对于使用元素为interface{}或者泛型类型的slice，应该怎么样进行初始化

A:

  官方对此没有好的办法，有下面比较复杂的实现方式：

```go
type Setter interface {
	Set(string)
}

// Settable is an integer type that can be set from a string.
type Settable int

func (p *Settable) Set(s string) {
	i, _ := strconv.Atoi(s)
	// 由于T是指针，所以指针值是nil，此处会panic
	*p = Settable(i)
}

// func (p *Settable) New() *Settable {
// 	return &Settable{}
// }

type Setter2[B any] interface {
	Set(string)
	*B // non-interface type constraint element
}

func FromStrings2[T any, PT Setter2[T]](s []string) []T {
	result := make([]T, len(s))
	for i, v := range s {
		// The type of &result[i] is *T which is in the type set
		// of Setter2, so we can convert it to PT.
		p := PT(&result[i])
		// PT has a Set method.
		p.Set(v)
	}
	return result
}

func main() {
	// nums := FromString[*Settable]([]string{"1", "2"})
	nums := FromStrings2[Settable]([]string{"1", "2"})
	fmt.Println(nums)
}
```