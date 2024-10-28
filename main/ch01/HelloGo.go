package main

// 使用某函数之前需要导入它的包
import (
	"fmt"
	"math"
)

// 不能导入无用的包，同样也不能声明未使用的变量
// import "strings"

func main() {
	// 表示函数是包的一部分
	fmt.Println("Hello Go")

	// go 语言中rune（符文）使用 Unicode 标准存储
	fmt.Println('a')

	fmt.Println(math.Floor(3.77))

	// 声明变量
	var a, b = "1", 2
	fmt.Println(a)
	fmt.Println(b)

	// 短变量声明，起初就知道它的默认值是多少，这样表示则更加明确
	c := 3
	fmt.Println(c)
}
