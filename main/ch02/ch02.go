package main

import (
	"strings"
	"time"
)

func main() {
	now := time.Now()
	println(now.Year())

	broken := "G# r#cks"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	print(fixed)
}
