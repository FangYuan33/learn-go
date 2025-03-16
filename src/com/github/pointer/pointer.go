package main

import "fmt"

func main() {
	s1, s2, _, e := multiReturn()
	if e == nil {
		fmt.Println(s1, s2)
	}
}

func multiReturn() (string, string, string, error) {
	return "1", "2", "2", nil
}

func double(a *int) {
	*a *= 2
}
