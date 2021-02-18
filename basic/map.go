package main

import (
	"fmt"
)

//trick with initialization before main func run

var globalMap = make(map[string]string)

func init() {
	globalMap["a"] = "b"
	globalMap["b"] = "a"
}

func main() {
	var mm = map[string]string{}
	mm["key"] = "value"
	fmt.Println(mm["key"])

	var mm1 = make(map[int]string)
	mm1[2] = "two"
	fmt.Println(mm1[2])

	fmt.Println(globalMap)
}
