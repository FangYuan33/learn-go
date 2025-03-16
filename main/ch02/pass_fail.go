package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	// 支持返回多个值，并且多个值必须被使用，如果想忽略可以使用 _ 占位
	//readString, _ := reader.ReadString('\n')
	//fmt.Println(readString)

	// go 语言中的变量命名遮盖导入的包就很烦
	//var fmt int = 1
	//log.Println(fmt)

	readString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	grade, err := strconv.ParseFloat(strings.TrimSpace(readString), 64)
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade == 100 {
		status = "Perfect!"
	} else if grade >= 60 {
		status = "Pass"
	} else {
		status = "Fail..."
	}
	log.Println("The grade of", grade, "is", status)

	a := 1
	b, c, a := 2, 3, 4

	log.Println(a)
	log.Println(b)
	log.Println(c)
}
