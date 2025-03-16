package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	processArray0(array)
	processArray(&array)
	fmt.Println(array)
}

func processArray0(array []int) {
	array[0] = 10
}

func processArray(array *[]int) {
	*array = append(*array, 6)
}
