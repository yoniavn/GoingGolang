/*
	Boolean type
	Numeric types
		-integers
		-floating point
		-complex numbers
	Text types
*/

package main

import "fmt"

func main() {
	//boolean types
	//** in go - every time we create a variable it initial to 0 - false in case of boolean
	var n bool //initially be false
	fmt.Printf("%v, %T\n", n, n)

	l := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", l, l)
	fmt.Printf("%v, %T\n", m, m)

	//numeric types
	/*
		int8 | uint8 | byte
		int16 | uint16
		int32 | uint32
		int64
	*/
	var num int16 = 33
	fmt.Printf("%v, %T\n", num, num)
}
