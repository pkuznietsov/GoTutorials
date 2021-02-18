package main

import "fmt"

func main() {
	var sl []int
	fmt.Println("empty slice", sl)

	sl = append(sl, 1, 2, 3, 4, 5)
	fmt.Println(sl)
}
