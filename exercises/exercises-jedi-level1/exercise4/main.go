package main

import (
	"fmt"
)

type bruno int

var x bruno

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)
}
