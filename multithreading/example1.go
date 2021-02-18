package main

import "fmt"

type myString string

func (*myString) process(i int) {
	fmt.Println(i)
}

func process(i int) {
	fmt.Println(i)
}

func main() {

	fmt.Println("start go routine")
	go process(1)

	go func() {
		fmt.Println("anon")
	}()

	for i := 0; i < 10; i++ {
		go process(i)
	}

	myStr := new(myString)
	go myStr.process(1)

	fmt.Scanln()
}
