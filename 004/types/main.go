package main

import "fmt"

var y = 42
var z = "Hi Fellas"
// Go is a static programming language
// So each variable is declared to hold a value of a certain type
// not a dynamic language

var x = `James said, "Shaken not stirred"`
func main()  {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	//z = 43 cannot do this as z is of type string
	z = "hello"
	fmt.Printf("z=%v is of type %T\n", z,z)
	fmt.Printf("x=%v is of type %T",x,x)
}

//Primitive: int, string, bool. Composite data types: struct, slice