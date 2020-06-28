package main

import (
	"fmt"
)

func main() {

	var i *int // declare i as an integer pointer. i is a memory address that holds an integer

	a := 100 // declare and assign a to be an integer, we will now store it at i

	i = &a // declare and assign i to the value 100

	fmt.Println(*i) // use the & operator to print out the memory address that I lives at, also know as a pointer. The type of &i is *int
}
