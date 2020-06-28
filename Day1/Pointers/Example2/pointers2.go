package main

import (
	"fmt"
)

func main() {

	i := 100 // declare and assign i to the value 100
	fmt.Println("Value of I:", i)

	fmt.Println("Address of I:", &i) // use the & operator to print out the memory address that I lives at, also know as a pointer. The type of &i is *int

	p := &i                                                     // store the address of i into p. & generates a pointer
	fmt.Println("Value of I by dereferencing its pointer:", *p) // print out the value that lives at the memory address of i, by dereferencing p, same thing as doing *&i or just printing i

	fmt.Println(*&i) // showcase that using the deference and
}
