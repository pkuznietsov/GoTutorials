package main

import "fmt"

func main() {

	//integers
	var i int = 2
	var autovar = -42
	var bigInt int64 = -123456789
	var unsignedBigInt uint64 = 123456789
	fmt.Println(i, autovar, bigInt, unsignedBigInt)

	//floats
	var f32 float32 = 3.14
	var f64 float64 = 4.30

	fmt.Println(f32, f64)

	//bools
	var bt bool = true
	var bf bool = false

	fmt.Println(bt, bf)

	//strings
	var hello string = "hello\n\t"
	var world string = "World"

	fmt.Println(hello, world)

	//Default values
	var empty_i int
	var empty_f float32
	var empty_string string
	var false_bool bool

	fmt.Println(empty_i, empty_f, empty_string, false_bool)

	//byte - alias for ascii
	var percentage byte = '\x25'
	fmt.Println(percentage)

	//short declare
	f42 := 42
	fmt.Println(f42)

	//cast
	fmt.Println(int(f42))

	//complex number
	complexNum := 2 + 3i
	fmt.Println(complexNum)

	//concat
	fmt.Println("a" + "b")

	//multiple def
	var (
		a int    = 3
		b string = "three"
	)
	fmt.Println(a, b)

}
