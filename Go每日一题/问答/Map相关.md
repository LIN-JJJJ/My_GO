如何确认两个 map 是否相等？

知识点：
1.  `==` `!=` 等比较运算符 ,不能用于 Slice, map, function 除了nil的情况,其余会编译报错
```text
//官方文档： https://go.dev/ref/spec#Comparison_operators
Slice, map, and function types are not comparable. However, as a special case, a slice, map, or function value may be compared to the predeclared identifier nil. Comparison of pointer, channel, and interface values to nil is also allowed and follows from the general rules above.
```
2、通过`reflect.DeepEqual`去判断

例：
> Slice
```go
package main

func main() {
	var a, b []int
	if a == b {
	}
}
//./main.go:5:5: invalid operation: a == b (slice can only be compared to nil)
```
```go
package main

import "fmt"

func main() {
	var a, b []int
	if a == nil {
		fmt.Println("a == nil T")
	}
	if b == nil {
		fmt.Println("a == nil T")
	}
}
//a == nil T
//a == nil T
```
>Map
```go
package main

func main() {
	var a, b map[int]int
	if a == b {
	}
}
//main.go:5:5: invalid operation: a == b (map can only be compared to nil)
```
```go
package main

import "fmt"

func main() {
	var a, b map[int]int
	if a == nil {
		fmt.Println("a == nil T")
	}
	if b == nil {
		fmt.Println("a == nil T")
	}
}
//a == nil T
//a == nil T
```
>function
```go
package main

func main() {
	var a, b func()
	if a == b {
	}
}
//./main.go:5:5: invalid operation: a == b (func can only be compared to nil)
```
```go
package main

import "fmt"

func main() {
	var a, b func()
	if a == nil {
		fmt.Println("a == nil T")
	}
	if b == nil {
		fmt.Println("a == nil T")
	}
}
//a == nil T
//a == nil T
```
> 数组 这是可以用 `==` `!=`
```go
package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3}
	var b = [...]int{1, 2, 3}
	if a == b {
		fmt.Println("T")
	}
}
//T
```
//因此这里比较map 需要用到反射
```text
func DeepEqual
用来判断两个值是否深度一致：除了类型相同；在可以时（主要是基本类型）会使用==；但还会比较array、slice的成员，map的键值对，结构体字段进行深入比对。map的键值对，对键只使用==，但值会继续往深层比对。DeepEqual函数可以正确处理循环的类型。函数类型只有都会nil时才相等；空切片不等于nil切片；还会考虑array、slice的长度、map键值对数。
```
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := make(map[string]int)
	n := make(map[string]int)
	m["1"] = 1
	n["1"] = 1

	fmt.Println(reflect.DeepEqual(m, n))
}
//true
```