package main

import (
	"fmt"
)

func main() {
	a := true
	if a {
		fmt.Println("True")
	}

	for {
		fmt.Println("first")
		break
	}

	i := 0
	for i < 10 {
		if i%2 == 0 {
			i++
			continue
		}
		fmt.Println(i)
		i++
	}

	index := 1
Loop:
	for {
		index++
		switch index {
		case 3:
			fmt.Println("break loop", index)
			break Loop
		}
	}

LoopOuter:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, j)
			if i == 0 && j == 1 {
				break LoopOuter
			}
		}
	}
}
