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

	//bitwise
	{
		a := 10 //1010
		b := 3  //0011
		fmt.Println(a & b)
		fmt.Println(a | b)
		fmt.Println(a ^ b)
		fmt.Println(a &^ b)

		shift := 8
		fmt.Println(shift << 3) //shift left 3 places
		fmt.Println(shift >> 3) //shift right 3 places
	}

	//Floating point numbers 32-bit and 64-bit
	{
		var zero float32 = 3.14 // expilicit declare float32
		one := 3.14
		one = 12.7e34
		one = 2.1e12 //float64
		fmt.Printf("%v, %T\n", zero, zero)
		fmt.Printf("%v, %T\n", one, one)

		//arithmatic operations
		h := 10.2
		j := 3.40
		fmt.Println("\narithmatic operations:")

		fmt.Println(h + j)
		fmt.Println(h - j)
		fmt.Println(h * j)
		fmt.Println(h / j)

	}

	//complex types
	{
		fmt.Println("\ncomplex types:")
		var com complex64 = 1 + 2i
		var com128 complex128 = 1 + 2i
		fmt.Printf("%v, %T\n", com, com)             //print the complex number
		fmt.Printf("%v, %T\n", com128, com128)       //print the complex number
		fmt.Printf("%v, %T\n", real(com), real(com)) //print the real part
		fmt.Printf("%v, %T\n", imag(com), imag(com)) //print the imaginary part
	}

	//string types  -represent ant UTF8 character
	{
		s := "\nthis is a string"
		fmt.Printf("%v, %T\n", s, s)

		s1 := "string can be acted like array s1[0] but needed to convert it with string(s[0])"
		fmt.Println(s1)
		fmt.Printf("%v, %T\n", string(s1[0]), s1[0])

		fmt.Println("strings are immutable - we cannot change values with s[0] = \"f\"")

		s2 := "S2 = string concatenation "
		s3 := "S3 = is possible"

		fmt.Printf("%v, %T\n", s2+s3, s2+s3)

		//represent the string as character array of bytes
		by := []byte(s3)
		fmt.Printf("%v, %T\n", by, by)
	}
}
