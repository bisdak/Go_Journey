package main

import "fmt"

func main() {
	// booleans
	var flag bool // no value assigned, set to false
	var isAwesome = true
	fmt.Println(flag)
	fmt.Println(isAwesome)

	// default value for numeric type is 0
	var i8 int8  
	var i16 int16
	fmt.Println(i8)
	fmt.Println(i16)

	//Strings
	//Immutable
	var text string // default value is empty string
	fmt.Println(text)
	text += "new_string"
	fmt.Println(text)

	//Type Conversions
	// var x int = 10
	// var y float64 = 30.2
	// var z float64 = float64(x) + y
	// var d int = x + int(y)
	// fmt.Println(z, d)

	//var Versus :=
	// use := inside function
	// var x int
	// var x, y int = 10, 20
	// var x, y int
	// var x, y = 10, "hello"
	// j := 10
	// k, l := 10, "hello"

	//Typed and Untyped Constants
	const typedX int = 10

	var y int = x
	var z float64 = x
	var d byte = x

	fmt.Println(y,z,d)
}