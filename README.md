## 从 Java 到 Go：面向对象的巨人与云原生的轻骑兵

Go 语言在 2009 年被 Google 推出，在创建之初便明确提出了“少即是多（Less is more）”的设计原则，强调“以工程效率为核心，用极简规则解决复杂问题”。它与 Java 语言生态不同，Go 通过编译为 **单一静态二进制文件实现快速启动和低内存开销**，**以25个关键字强制代码简洁性**，**用接口组合替代类继承**，**以显式返回error取代异常机制** 和 **轻量级并发模型（Goroutine/Channel）** 在 **云原生基础设施领域** 占据主导地位，它也是 Java 开发者探索云原生技术栈的关键补充。本文将对 Go 语言和 Java 语言在一些重要特性上进行对比，为 Java 开发者在阅读和学习 Go 语言相关技术时提供参考。

### 代码组织的基本单元

在 Java 中，我们会创建 `.java` 文件作为 **类**（类名与文件名相同），并在该类中定义相关的字段或方法等（OOP），如下定义 `User` 和 `Address` 相关的内容便需要声明两个 `.java` 文件（`User.java`, `Address.java`）定义类：

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

而在 Go 语言中，它是通过 **“包”** 来组织代码的：每个目录下的所有 `.go` 文件共享同一个 **包**，在包内可以定义多个结构体、接口、函数或变量。它并不要求文件名与声明的内容一致，比如创建 `User` “结构体”并不会要求 `.go` 文件也命名为 `User.go`，而是任何命名都可以（命名为 `user.go` 甚至 `a.go` 这种无意义的命名），而且同一个包下可以创建多个 `.go` 文件。如下为在 `user` 包下定义 `User` 和 `Address` 相关的内容，它们都被声明在一个 `user.go` 文件中：

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

相比来说，Java 代码组织的基本单元是类，作为面向对象的语言更侧重对象定义，而 Go 代码组织的基本单元是包，并更侧重功能模块的聚合。

#### 可见性控制

在 Java 中通过 `public/protected/private` 关键字控制成员的可见性，而在 Go 语言中，通过 **首字母大小写** 控制“包级别的导出”（大写字母开头为 `public`），**包的导出成员对其他包可见**。以 user 包下 `User` 类型的定义为例，在 main 包下测试可见性如下：

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

在 Java 语言中，对变量（静态变量或局部变量）的声明只有一种方式，“采用 = 运算符赋值”显式声明（在 Jdk 10+支持 var 变量声明），如下：

```java
public class Test {
    public static void main(String[] args) {
        int x = 100;
    }
}
```

而在 Go 语言中，变量声明有两种主要方式：**短声明（`:=` 运算符）** 和 **长声明（`var` 声明）**，它们的适用场景和限制有所不同，以下是详细区分：

#### 短声明（`:=`）

只能在函数（包括 `main`、自定义函数或方法、`if/for` 块等）内部使用，不能在包级别（全局作用域）使用，并且 **声明的局部变量必须被使用**，不被使用的局部变量不能被声明：

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

这种短声明直接根据右侧值自动推断变量类型，无需显式指定类型，并且可以一次性声明多个变量，但至少有一个变量是 **新声明的**：

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

#### 常量的声明

Go 语言中对常量的声明采用 `const` 关键字，并且在声明时便需要被赋值，如下所示：

```go
package main

import "fmt"

// DaysInWeek const 变量名 类型 = 具体的值
const DaysInWeek int = 7

func main() {
   const name = "abc"
   fmt.Println(name, DaysInWeek)
}
```

在 Java 语言中对常量的声明会使用 `static final` 引用：

```java
public class Constants {
    public static final int DAYS_IN_WEEK = 7;
    
    // ...
}
```

### 方法/函数的声明

在 Go 语言中，方法的声明遵循 **func (接收器) 方法名(入参) 返回值** 的格式，无返回值可以不写（无需 void 声明），通过 **接收器（Receiver）** 将方法绑定到结构体上，如下为 `User` 结构体方法的声明：

```go
package user

type User struct {
	name string
}

// Name (u *User) 即为接收器，表示该方法绑定在了 User 类型上
func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}
```

而“函数”的声明不需要定义接收器，遵循的是 **func 方法名(入参) 返回值** 的格式。Go 语言中的函数类似于 Java 语言中的静态方法，以下是声明将整数扩大两倍的函数：

```go
package main

func double(a *int) {
	*a *= 2
}
```

并且，在 Go 语言中，方法/函数支持多返回值（常用于错误处理），并且如果并不需要全部的返回值，可以用 `_` 对返回值进行忽略，因为Go语言不允许定义未使用的局部变量，如下所示：

```go
package main

import "fmt"

func main() {
	// 忽略掉了第三个返回值
	s1, s2, _, e := multiReturn()
	if e == nil {
		fmt.Println(s1, s2)
	}
}

func multiReturn() (string, string, string, error) {
	return "1", "2", "2", nil
}
```

此外，接收器参数和函数的形参支持传入指针，用 `*` 符号表示。在 Go 语言中有指针的概念，我们在这里说明一下：Go 语言是 **“值传递”** 语言，方法/函数的形参（或接收器）如果不标记指针的话，接收的实际上都是 **实参的副本**，那么 **在方法/函数中的操作并不会对原对象有影响**。如果想对原对象进行操作，便需要通过指针获取到原对象才行（因为值传递会对原对象和形参对象都划分空间，所以针对较大的对象都推荐使用指针以节省内存空间）。在如下示例中，如果我们将上文中 `double` 方法的形参修改为值传递，这样是不能将变量 a 扩大为两倍的，因为它操作的是 a 变量的副本：

```go
package main

import "fmt"

func main() {
	a := 5
	double(a)
	// 想要获取 10，但打印 5
	fmt.Println(a)
}

func double(a int) {
	a *= 2
}
```

想要实现对原对象 a 的操作，便需要使用指针操作，将方法的声明中传入指针变量 `*int`：

```go
package main

import "fmt"

func main() {
	a := 5
	// & 为取址运算符
	double(&a)
	// 想要获取 10，实际获取 10
	fmt.Println(a)
}

// *int 表示形参 a 传入的是指针
func double(a *int) {
	// *a 表示从地址中获取变量 a 的值
	*a *= 2
}
```

再回到 `User` 类型的声明中，如果我们将接收器修改成 `User`，那么 `SetName` 方法是不会对原变量进行修改的，它的修改实际上只针对的是 `User` 的副本：

```go
package user

type User struct {
	name string
}

// SetName 指定为值接收器
func (u User) SetName(name string) {
	u.name = name
}
```

这样 `SetName` 方法便不会修改原对象，`SetName` 的操作也仅仅对副本生效了：

```go
package main

import (
	"fmt"
	"learn-go/src/com/github/user"
)

func main() {
	u := user.User{}
	u.SetName("abc")
	// 实际输出为 {}，并没有对原对象的 name 字段完成赋值
	fmt.Println(u)
}
```

在 Java 中并没有指针的概念，Java 中除了基本数据类型是值传递外，其他类型在方法间传递的都是“引用”，对引用对象的修改也是对原对象的修改。

### 接口

Go 语言也支持接口的声明，不过相比于 Java 语言它更追求 **“灵活与简洁”**。Go 的接口实现是“隐式地”，**只要类型实现了接口的所有方法，就自动满足该接口**，无需显式声明。如下：

```go
package writer

type Writer interface {
   Write([]byte) (int, error)
}

// File 无需声明实现 Writer，实现了接口所有的方法便自动实现了该接口
type File struct{}

func (f *File) Write(data []byte) (int, error) {
   return len(data), nil
}
```

Java 语言则必须通过 `implements` 关键字声明类对接口的实现：

```java
public interface Writer {
   int write(byte[] data);
}

public class File implements Writer {  // 必须显式声明
   @Override
   public int write(byte[] data) {
      return data.length;
   }
}
```

它们对类型的判断也是不同的，在 Go 语言中采用如下语法：

```go
package writer

func typeTransfer() {
   var w Writer = File{}
   // 判断是否为 File 类型，如果是的话 ok 为 true
   f, ok := w.(File)
   if ok {
      f.Write(data)
   }
}
```

而在 Java 语言中则采用 `instanceof` 和强制类型转换：

```java
private void typeTransfer() {
   Writer w = new File();
   if (w instanceof File) {
      File f = (File) w;
      f.write(data);
   }
}
```

Go 语言还采用空接口 `interface{}` 来表示任意类型，作为方法入参时则支持任意类型方法的传入，类似 Java 中的 `Object` 类型：

```go
package writer

func ProcessData(data interface{}) {
	// ...
}
```

除此之外，Go 语言在 1.18+ 版本引入了泛型，采用 `[T any]` 方括号语法定义类型约束，`any` 表示任意类型，如果采用具体类型限制则如下所示：

```go
package writer

// Stringer 定义约束：要求类型支持 String() 方法
type Stringer interface {
    String() string
}

func ToString[T Stringer](v T) string {
    return v.String()
}
```

通过类型的限制便能使用类型安全替代空接口 `interface{}`，避免运行时类型断言：

```go
// 旧方案：空接口 + 类型断言
func OldMax(a, b interface{}) interface{} {
    // 需要手动断言类型，易出错
}

// 新方案：泛型
func NewMax[T Ordered](a, b T) T { /* 直接比较 */ }
```

泛型还在通用数据结构上有广泛的应用：

```go
type Stack[T any] struct {
    items []T
}
func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}
```

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

---

### “引用类型”

在 Go 语言中，**严格来说并没有“引用类型”这一官方术语**，但在 Go 语言社区中通常将 **Slice（切片）、Map（映射）、Channel（通道）** 称为“引用语义类型”（或简称引用类型），因为它们的行为与传统的引用类型相似，在未被初始化时为 `nil`，并无特定的“零值”。除了这三种类型之外，Go 的其他类型（如结构体、数组、基本类型等）都是 **值类型**。

#### Slice 

Go 的 **Slice** 本质上是动态数组的抽象，基于底层数组实现自动扩容。它类似于 Java 中的 `ArrayList`，采用 `var s []int` 或 `s := make([]int, 5)` 声明，如下：

```go
package main

import "fmt"

func slice() {
  // 初始化到小为 0 的切片
  s := make([]int, 0)
  // 动态追加元素
  s = append(s, 1, 2, 3, 4, 5)
  fmt.Println(s)
  // 子切片，左闭右开区间 sub = {2, 3}
  sub := s[1:3]
  fmt.Println(sub)
  // 修改子切片值会影响到 s 原数组
  sub[0] = 99
  fmt.Println(s)
}
```

切片的底层数组并不能增长大小。如果数组没有足够的空间来保存新的元素，所有的元素会被拷贝至一个新的更大的数组，并且切片会被更新为引用这个新的数组。但是由于这些场景都发生在 `append` 函数内部，所发知道返回的切片和传入 `append` 函数的切片是否为相同的底层数组，所以如果保留了两个切片，那么这一点需要注意。

#### Map

Go 的 Map 本质上是无序键值对集合，基于哈希表实现。它的键必须支持 `==` 操作（如基本类型、结构体、指针），声明方式为 `m := make(map[string]int)` 或 `m := map[string]int{"a": 1}`，它与 Java 中的 `HashMap` 类似，如下所示：

```go
package main

import "fmt"

func learnMap() {
  m := make(map[string]int)
  m["a"] = 1
  // 安全的读取
  value, ok := m["a"]
  if ok {
    fmt.Println(value)
  }
  delete(m, "a")
}
```

#### Channel

Go 的 Channel 是用于 **协程（goroutine，Go 语言中的并发任务类似 Java 中的线程）间通信** 的管道，支持同步或异步数据传输。无缓冲区通道会阻塞发送/接收操作，直到另一端就绪。它的声明方式为 `channel := make(chan string)`（无缓冲）或 `channel := make(chan string, 3)`（有缓冲，缓冲区大小为 3），创建无缓存区的 channel 示例如下：

```go
package main

import "fmt"

// 创建没有缓冲区的 channel，如果向其中写入值后而没有其他协程从中取值，
// 再向其写入值的操作则会被阻塞，也就是说“发送操作会阻塞发送 goroutine，直到另一个 goroutine 在同一 channel 上执行了接收操作”
// 反之亦然
func channel() {
  channel1 := make(chan string)
  channel2 := make(chan string)

  // 启动一个协程很简单，即 go 关键字和要调用的函数
  go abc(channel1)
  go def(channel2)

  // <- 标识符指出 channel 表示从协程中取值，输出一直都会是 adbecf
  fmt.Print(<-channel1)
  fmt.Print(<-channel2)
  fmt.Print(<-channel1)
  fmt.Print(<-channel2)
  fmt.Print(<-channel1)
  fmt.Println(<-channel2)
}

// <- 标识符指向 channel 表示向 channel 中发送值
func abc(channel chan string) {
  channel <- "a"
  channel <- "b"
  channel <- "c"
}

func def(channel chan string) {
  channel <- "d"
  channel <- "e"
  channel <- "f"
}
```

如果创建有缓冲的 channel，在我们的例子中，那么就可以实现写入协程不必等待 main 协程的接收操作了：

```go
package main

import "fmt"

func channelNoBlocked() {
	// 表示创建缓冲区大小为 3 的 channel，并且 channel 传递的类型为 string
	channel1 := make(chan string, 3)
	channel2 := make(chan string, 3)

	go abc(channel1)
	go def(channel2)

	// 输出一直都会是 adbecf
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Println(<-channel2)
}
```

> 在 Go 中创建上述三种引用类型的对象时，都使用了 `make` 函数，它是专门用于初始化这三种引用类型的，如果不使用该函数，直接声明（如`var m map[string]int`）会得到 `nil` 值，而无法直接操作。它与 Java 中的 `new` 关键字操作有很大的区别，`new` 关键字会为对象分配内存 **并调用构造函数**（初始化逻辑在构造函数中），而在 Go 的设计中是没有构造函数的，Go 语言除了这三种引用类型，均为值类型，直接声明即可，声明时便会直接分配内存并初始化为零值。

### 从失败中恢复

在 Go 语言中 **没有传统“异常”概念**，它不依赖 `try/catch`，而是通过 **显式返回错误值** 和 `panic/recover` 机制处理。它的错误（error）也是普通的数据，能够作为值传递。在多数方法中能看到如下类似的实现：

```go
package main

func main() {
	data, err := ReadFile("file.txt")
	// 处理错误
	if err != nil {
		log.Fatal(err)
	}
	// ...
}

func ReadFile(path string) ([]byte, error) {
    // 成功返回 data, nil
    // 失败返回 nil, error
}
```

Go 语言使用 `panic` 来处理不可恢复的或程序无法继续运行的错误（如数组越界、空指针），这类似于 Java 语言中的 `throw` 异常，它会中断方法或函数的执行，向上抛出直到遇到 `defer` 和 `recover()` 函数的声明捕获或者程序崩溃： 

```go
// 初始化失败时触发 panic
func initDatabase() {
    if !checkDatabaseConnection() {
        panic("Database connection failed!")
    }
}

// 通过 recover 捕获 panic
func main() {
	// 延迟函数的执行
    defer func() {
		// 使用 recover() 函数尝试捕获异常 
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    initDatabase()
    // 正常逻辑...
}
```

`defer` 关键字 **必须修饰的函数或方法**，而且被这个关键字修饰的函数或方法 **一旦注册** 无论如何都会被执行（类似于 Java 中的 `finally`），但如果 `defer` 声明在函数尾部，但函数在运行到该 `defer` 语句之前就退出（例如中途 `return` 或 `panic`），则 `defer` **不会注册，也不会执行**。**所以该关键字在资源被初始化之后应该立即使用，而非像 Java 一样声明在方法的尾部**。而且 `defer` 支持声明多个，但执行的顺序是逆序的。

`revocer()` 函数与 `defer` 关键字搭配使用，它会返回函数执行过程中抛出的 `panic`（未发生 `panic` 时会为 `nil`），可以帮助开发者恢复或提供有用的异常信息。

以下是在文件读取场景 Go 和 Java 语言在语法上的不同：

- Go

```go
func readFile() {
    file, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    // 处理文件内容
}
```

- Java

```java
public void readFile() {
    // try-with-resources
    try (FileReader file = new FileReader("file.txt")) {
        // 处理文件内容
    } catch (IOException e) {
        System.err.println("Error: " + e.getMessage());
    }
}
```

> 问：我看到其他编程语言有 `exception`。`panic` 和 `recover` 函数似乎以类似的方式工作。我可以把它们当作 `exception` 来使用吗？
> 
> 答：Go语言维护者强烈建议不要这样做。甚至可以说，语言本身的设计不鼓励使用 `panic`和 `recover`。在 2012 年的一次主题会议上，RobPike（Go的创始人之一）把 `panic` 和 `recover` 描述为“故意笨拙”。这意味着，在设计 Go 时，创作者们没有试图使 `panic` 和 `recover` 被容易或愉快地使用，因此它们会很少使用。这是 Go 设计者对 `exception` 的一个主要弱点的回应：它们可以使程序流程更加复杂。相反，Go 开发人员被鼓励以处理程序其他部分的方式处理错误：使用 `if` 和 `return` 语句，以及 `error` 值。当然，直接在函数中处理错误会使函数的代码变长，但这比根本不处理错误要好得多。（Go的创始人发现，许多使用 `exception` 的开发人员只是抛出一个 `exception`，之后并没有正确地处理它。）直接处理错误也使错误的处理方式一目了然，你不必查找程序的其他部分来查看错误处理代码。所以不要在 Go 中寻找等同于 `exception` 的东西。这个特性被故意省略了。对于习惯了使用 `exception` 的开发人员来说，可能需要一段时间的调整，但 Go 的维护者相信，这最终会使软件变得更好。

### for 和 if

#### for

Go 语言的循环语法只有 `for`，没有 `while` 或 `do-while`，但可实现所有循环模式：

```go
// 1. 经典三段式（类似 Java 的 for 循环）
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// 2. 类似 while 循环（条件在前）
sum := 0
for sum < 10 {
    sum += 2
}

// 3. 无限循环（省略条件）
for {
    fmt.Println("Infinite loop")
    break  // 需手动退出
}

// 4. 遍历集合，采用 range 关键字，index 和 value 分别表示索引和值
arr := []int{1, 2, 3}
for index, value := range arr {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

#### if

Go 语言的 `if` 语法相比于 Java 支持声明 + 条件的形式，并且强制要求大括号（即使是单行语句也必须使用 `{}`）：

```go
// 支持简短声明（声明 + 条件）
if num := 10; num > 5 {  
    fmt.Println("num is greater than 5")
}
// 简单判断
if num > 5 {
    fmt.Println("num is greater than 5")
}
```

---

### 巨人的肩膀

- 《Head First Go 语言程序设计》