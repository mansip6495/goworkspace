package main

import "fmt"

var x int
var y string
var z bool
type made int
var a made
var b int

func main()  {

	// exercise1()
	// exercise2()
	// exercise3()
	// exercise4()
	// exercise5()
	
}

func exercise1()  {

	// Shorthand declaration
	x := 42
	y := "James Bond"
	z := true

	// Single print statement
	fmt.Printf("\nSingle Print statement: x = %v, y = %v, z=%v\n",x,y,z)
	fmt.Println("Single Print Statement",x,y,z)

	// Multiple print statement
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func exercise2()  {
	// Print zero value of package level variables
	fmt.Printf("\n x = %T/%v, y = %T/%v, z = %T/%v", x, x, y, y, z, z)
}

func exercise3()  {
	x= 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("%v %v %v",x,y,z)
	fmt.Println(s)
}

func exercise4() {
	fmt.Printf("a = %v, %T\n", a, a)
	a = 42
	fmt.Println(a)
}

func exercise5()  {
	fmt.Printf("a = %v, %T\n", a, a)
	a = 42
	fmt.Println(a)
	b = int(a)
	fmt.Printf("b = %v, %T", b, b)
}

func exercise6()  {
	
}