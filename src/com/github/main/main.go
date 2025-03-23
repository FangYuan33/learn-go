package main

import (
	"fmt"
	"learn-go/src/com/github/user"
)

var global int = 1

func main() {
	var u user.User
	// fmt.Println(u.name)
	fmt.Println(u.Name())

	//slice()
	//learnMap()
	//channel()
	channelNoBlocked()
}

func slice() {
	// 初始化到小为 0 的切片
	s := make([]int, 0)
	// 动态追加元素
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s)
	// 子切片，左闭右开区间 sub {2, 3}
	sub := s[1:3]
	fmt.Println(sub)
	// 修改子切片值会影响到 s 原数组
	sub[0] = 99
	fmt.Println(s)
}

func learnMap() {
	m := make(map[string]int)
	m["a"] = 1
	value, ok := m["a"]
	if ok {
		fmt.Println(value)
	}
	delete(m, "a")
}

// 创建没有缓冲区的 channel，如果向其中写入值后而没有其他协程从中取值
// 再向其写入值的操作则会被阻塞，也就是说“发送操作会阻塞发送 goroutine，
// 直到另一个 goroutine 在同一 channel 上执行了接收操作”，反之亦然
func channel() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	// 启动一个协程很简单，即 go 关键字和要调用的函数
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
