// this example shows the difference between passing pointers and values to functions
package main

import "fmt"

type Point struct {
	x int
	y int
}

// under the hood, golang copies the point made in main, and modifies the copy in this function
// without returning the modified copy point, the origional is not modified
func ChangePointByValue(p Point) {
	p.x = 0
	p.y = 0
}

// here, we pass this function a pointer to the point created in the main function
// thus, it gets changed, and no copying or returning is required
func ChangePointByReference(p *Point) {
	p.x = 0
	p.y = 0
}

func main() {

	p := Point{x: 50, y: 25}
	fmt.Println("Point Before Modification:", p)

	ChangePointByValue(p)
	fmt.Println("Point After Modification by value:", p)

	ChangePointByReference(&p) // pass by reference (remember & signifies the memory address of the variable)
	fmt.Println("Point After Modification by reference:", p)

}

// General rules of thumb of when to pass by pointer
// if variable is expected to be modified, it must be passed by pointer.
// if a variable is very large (a large struct, for example), its better to pass that variable by a pointer to avoid copying the large variable in memory
// if the variable is a map or slice, it should be passed by value, since they are pointers by default
