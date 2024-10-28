## Go 语言与 Java 语言的不同

### Hello Go

Go 语言不能导入未使用到的包，并且函数是基于包的一部分，而并不是像 Java 一样基于类。比如 `fmt.Println` 方法，这个方法是在 `fmt` 包下的，调用时也是以包名为前缀。

```go
package main

// 使用某函数之前需要导入它的包
import (
	"fmt"
	// 不能导入未使用到的包
	//"math"
)

func main() {
	// 表示函数是包的一部分
	fmt.Println("Hello Go")
}
```

### 字符

Go 语言的字符（rune）使用 Unicode 来存储，而并不是字符本身，如果把 rune 传递给 `fmt.Println` 方法，会在控制台看到数字。虽然 Java 语言同样以 Unicode 保存字符（char），不过它会在控制台打印字符信息。

```go
package main

import "fmt"

func main() {
	// 97
	fmt.Println('a')
}
```

### 声明变量

在 go 语言中，声明变量有两种方式，如下：

```go
package main

import "fmt"

func main() {
	// 1.
	var a int = 1
	// 注意，未被使用到的变量是不能够被声明的
	fmt.Println(a)
	
	// 当然它也能将类型标注舍去
	var b = 2
	fmt.Println(b)
	// 也能声明多个，而且还能指定不同的类型
	var c, d, e = 3, 4, 1.1
	fmt.Println(c, d, e)
	
	// 2. := 短声明。如果起初就知道某变量的值，短声明会更加常用
	f, g := 2, 3
	fmt.Println(f, g)
}
```

### 大小写区分 public 和 private

- 如果变量、函数或类型的名称以 **大写字母开头**，则认为它是导出的，可以从当前包之外的包访问（对应上文中 `fmt.Println` 方法，开头的 P 字母便是大写的）
- 如果变量、函数或类型的名称以 **小写字母开头**，则认为它是未导出的，只能在当前包访问

而在 Java 语言中，定义了 `public` 和 `private` 来表示变量、函数和类型是公开的还是类内私有的。

### 类型转换

Go 语言类型转换与 Java 不同，并且它不支持不相同类型的变量操作，如下：

```go
package main

import "fmt"

func main() {
	a := 1
	b := 2.2
	// 如果不类型转换则不能通过编译
	fmt.Println(float64(a) * b)
}
```

在 Java 语言中，不同类型的变量 `a` 和变量 `b` 会被转换成 `float` 类型计算并得出结果。