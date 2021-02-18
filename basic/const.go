package main

import "fmt"

const name int = 137
const myname string = "pavel"

//const group
const (
	flag1 = 1
	flag2 = 2
)

const (
	zerovar = iota
	onevar
	_
	threevar
)

const (
	z1 = iota + 1
	z2
	z3
)

func main() {
	fmt.Println(zerovar, onevar, threevar)
	fmt.Println(z1, z2, z3)
}
