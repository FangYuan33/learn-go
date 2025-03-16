## 从 Java 到 Go：面向对象的巨人与云原生的轻骑兵

Go 语言在 2009 年被 Google 推出，在创建之初便明确提出了“少即是多（Less is more）”的设计原则，强调“以工程效率为核心，用极简规则解决复杂问题”。它与 Java 语言生态不同，Go 通过编译为 **单一静态二进制文件实现快速启动和低内存开销**，**以25个关键字强制代码简洁性**，**用接口组合替代类继承**，**以显式返回error取代异常机制** 和 **轻量级并发模型（Goroutine/Channel）** 在 **云原生基础设施领域** 占据主导地位，它也是 Java 开发者探索云原生技术栈的关键补充。本文将对 Go 语言和 Java 语言在一些重要特性上进行对比，为大家以后阅读和学习 Go 语言相关技术资料提供参考。

### 代码组织的基本单元

在 Java 中，每个 `.java` 文件必须包含与文件名相同的 **类**（public 类），并在该类中定义相关的字段或方法等（OOP），如下定义 User 和 Address 相关的内容便需要声明两个 `.java` 文件定义类：

```java
public class User {

    private String name;
    
    public String getName() {
        return name;
    }
    public void setName(String name) {
        this.name = name;
    }
}
```

```java
public class Address {
    private String city;
    
    public String getCity() {
        return city;
    }
    public void setCity(String city) {
        this.city = city;
    }
}
```

而在 Go 语言中，每个目录下的所有 `.go` 文件共享同一个 **包名**，在包内可以定义多个类型、接口、函数和变量，如下为在 `user` 包下定义 User 和 Address 相关的内容：

```go
package user

type User struct {
   name string
}

func (u *User) Name() string {
   return u.name
}

func (u *User) SetName(name string) {
   u.name = name
}

type Address struct {
   city string
}

func (a *Address) City() string {
   return a.city
}

func (a *Address) SetCity(city string) {
   a.city = city
}
```

相比来说，Java 的类更侧重对象定义，而 Go 的包更侧重功能模块的聚合。

#### 可见性控制

在 Java 中是通过 `public/protected/private` 修饰符控制成员可见性，而在 Go 语言中，通过 **首字母大小写** 控制“导出”（大写字母开头为 `public`），包的导出成员对其他包可见。以 user 包下 User 类型的定义为例，在 main 包下测试可见性如下：

```go
package main

import (
	"fmt"
	// user package 的全路径
	"learn-go/src/com/github/user"
   // 不能导入未使用到的包
   //"math"
)

func main() {
	var u user.User
	// 在这里是不能访问未导出的字段 name
	// fmt.Println(u.name)
	fmt.Println(u.Name())
}
```

Go 语言不能导入未使用到的包，并且函数是基于包的一部分。比如 `fmt.Println` 函数，这个函数是在 `fmt` 包下的，调用时也是以包名为前缀。

### 变量的声明

在 Java 语言中，对变量（静态变量或局部变量）的声明只有一种方式，“采用 = 运算符赋值”进行显式声明（在 Jdk 10+支持 var 变量声明），如下：

```java
public class Test {
    public static void main(String[] args) {
        int x = 100;
    }
}
```

而在 Go 语言中，变量声明有两种主要方式：**长声明（`var` 声明）** 和 **短声明（`:=` 运算符）**，它们的适用场景和限制有所不同，以下是详细区分：

#### 短声明（`:=`）

只能在函数（包括 `main`、自定义函数、`if/for` 块等）内部使用，不能在包级别（全局作用域）使用，并且 **声明的局部变量必须被使用**，不被使用的局部变量不能被声明：

```go
package main

import "fmt"

func main() {
	// 正确
	x := 10
	fmt.Println(x)
	// 未被使用，不能被声明
	// y := 20
	// 不赋值也不能被声明
	// z :=            
}

// 错误：不能在包级别使用短声明
// y := 20          
```

这种短声明直接根据右侧值自动推断变量类型，无需显式指定类型，并且可以一次性声明多个变量，但至少有一个变量是**新声明的**：

```go
package main

import "fmt"

func main() {
	// 同时声明 a 和 b
	a, b := 1, "abc"
	// c 是新变量，b 被重新赋值
	c, b := 2, "def"
	// 无新变量无法再次对已声明的变量再次声明
	//a, b := 4, "error"
	
	fmt.Println(a, b, c)
}
```

#### 长声明（`var` 声明）

在全局作用域声明变量必须使用 `var`；在需要延迟初始化时也需要采用长声明；显示指定类型也需要使用长声明

```go
package main

import "fmt"

var global int = 42

func main() {
	// a = 0
	var a int
	// s = ""
	var s string
	// 未被初始化值会默认为“零”值，a 为 0，s 为空字符串
	fmt.Println(a, s)
}
```

函数内部的局部变量，尤其是需要类型推断和简洁代码时优先用短声明；在包级别声明变量，需要显式指定类型或声明变量但不立即赋值（零值初始化）时，使用长声明。

在 Go 语言中还有一点需要注意：**声明变量时，应确保它与任何现有的函数、包、类型或其他变量的名称不同**。如果在封闭范围内存在同名的东西，变量将对它进行覆盖，也就是说，优先于它，如下所示：

```go
package main

import "fmt"

func main() {
    // 这个变量会把导入的 fmt 包覆盖掉
	fmt := 1
	println(fmt)
}
```

那么我们导入的 `fmt` 包在被局部变量覆盖后便不能再被使用了。

### 基本数据类型

Go 的基本数据类型分为 **4 大类**，相比于 Java 更简洁且明确：

| 类别      | 具体类型                                                     | 说明                                                                                                                            |
|---------|----------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------|
| **数值型** | `int`, `int8`, `int16`, `int32`, `int64`                 | Go 的 `int` 长度由平台决定（32 位系统为 4 字节，64 位为 8 字节），有符号整数（位数明确，如 `int8` 占 1 字节）                                                       |
|         | `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` | 无符号整数（`uintptr` 用于指针运算）                                                                                                       |
|         | `float32`, `float64`                                     | 浮点数（默认 `float64`）                                                                                                             |
|         | `complex64`, `complex128`                                | 复数（实部和虚部分别为 `float32` 或 `float64`，Java 无此类型）                                                                                  |
| **布尔型** | `bool`                                                   | 仅 `true`/`false`（不可用 0/1 替代）                                                                                                  |
| **字符串** | `string`                                                 | 不可变的 UTF-8 字符序列                                                                                                               |
| **派生型** | `byte`（=`uint8`）                                         | 1 字节数据                                                                                                                        |
|         | `rune`（=`int32`）                                         | Go 语言的字符（rune）使用 Unicode 来存储，而并不是字符本身，如果把 rune 传递给 `fmt.Println` 方法，会在控制台看到数字。虽然 Java 语言同样以 Unicode 保存字符（char），不过它会在控制台打印字符信息 |

Go 和 Java 同样都是 **静态类型语言**，要求在 **编译期** 确定所有变量的类型，且类型不可在运行时动态改变。Go 不允许任何隐式类型转换（如 `int32` 到 `int64`），但是在 Java 中允许基本类型隐式转换（如 `int` → `long`），除此之外，Go 语言会严格区分类型别名（如 `int` 与 `int32` 不兼容）。在 Go 语言中如果需要将不同类型的变量进行计算，需要进行类型转换：

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






| **特性**   | **Go**                                  | **Java**              |
|----------|-----------------------------------------|-----------------------|
| **接口实现** | 隐式实现（只需实现接口方法，无需 `implements`）          | 显式声明 `implements` 接口  |
| **空接口**  | `interface{}` 可表示任意类型（类似 Java `Object`） | `Object` 是根类，但需强制类型转换 |

**Go 的隐式接口示例**：
```go
package main

type Writer interface { Write([]byte) error }

type MyWriter struct {}

func (w MyWriter) Write(data []byte) error { return nil }
// MyWriter 自动实现 Writer 接口，无需声明
```

**Java 的显式接口示例**：

```java
interface Writer { void write(byte[] data); }

class MyWriter implements Writer { 
    @Override 
    public void write(byte[] data) {} 
}
// 必须显式声明 implements
```

---

### 巨人的肩膀

- 《Head First Go 语言程序设计》