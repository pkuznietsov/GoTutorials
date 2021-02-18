package main

import "fmt"

const size uint = 5

func main() {
	var a1 [10]int
	fmt.Println(len(a1))

	var a2 [5 * size]bool
	fmt.Println(len(a2))

	var a3 = [...]int{3, 3, 3}
	fmt.Println(len(a3))

	var a [3][3]int
	a[1][1] = 3
	fmt.Println(a[1][1])
}
