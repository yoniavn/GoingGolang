package main

import (
	"fmt"
)

/*
	-naming conventions
	-typed constants
	-untyped constants
	-enumerated constants
	-enumeration expressions
*/

//decleration at the package level
const a int16 = 27

func main() {
	//declare constants
	const myConst int = 40 //must be assign at compile time

	fmt.Printf("%v, %T\n", myConst, myConst)

	const a int = 22 //the decleration at the main level was won over the package level  
	fmt.Printf("%v, %T\n", a, a)

}
