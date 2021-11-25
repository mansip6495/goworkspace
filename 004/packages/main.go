package main

import "fmt"

func main(){
	// func Println(a ...interface{}) (n int, err error)
	// ... variadic parameter: any number of any type of values. Interface has no type
	// ... means any number of
	// package is an ready code you can use. Imports bring it in your code
	n,e := fmt.Println("Hello", 42, true)
	fmt.Println(n, e)
}