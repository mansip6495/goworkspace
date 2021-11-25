package main

import "fmt"

var y = 43

var z int
// declares a variable z with identifier z of type int and assigns value 0 of type int

func main()  {
	// Declare a variable and assign a value
	// short declaration operator := 
	x := 42
	fmt.Println(x)
	
	fmt.Println(y)
	// var is needed to declare a variable outside function
	// so use short declaration in function and "var" outside function  
}