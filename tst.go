package main

import (
	"fmt"
)

func main() {
	var cache map[string]string
	cache = make(map[string]string)
	var u string = "1"
	_, ok := cache[u]
	fmt.Println(ok)
}
