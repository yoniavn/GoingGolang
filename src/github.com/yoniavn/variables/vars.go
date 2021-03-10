package main

import (
	"fmt"
	"strconv"
)

//--------package-level scope

var I int = 4

//declare block of variables
var (
	one int    = 1
	two int    = 2
	str string = "string var"
)

//--------end package-level scope

func main() {
	//declare variables
	var i int = 4
	fmt.Println(i)

	//let compiler decide the type
	ii := 3
	fmt.Println(ii)

	//format the string to display value and type
	fmt.Printf("value: %v, type: %T\n", ii, ii)

	var f float32 = 4.3
	fmt.Printf("value: %v, type: %T\n", f, f)

	//conversion to float
	var jj float32
	jj = float32(i)
	fmt.Printf("value: %v, type: %T\n", jj, jj)

	//convert int to string using strconv package
	var number int = 40
	var str string = strconv.Itoa(number)
	fmt.Printf("value: %v, type: %T", str, str)

}
