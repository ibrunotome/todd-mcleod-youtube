package main

import (
	"fmt"
)

type bruno int

var x bruno

var y int

func main() {
	fmt.Println(x)
	fmt.Println("%T\n", x)

	x = 21

	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\t", y)
}
