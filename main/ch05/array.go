package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(array)

	sum := 0
	// 想忽略其中的值可以使用 _ 占位
	for index, value := range array {
		fmt.Println(index, value)
		sum += value
	}
	fmt.Printf("average: %.2f \n", float64(sum/len(array)))
}
