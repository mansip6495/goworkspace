package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64
var z int8 = -128

func main()  {
	// x := 42
	// y := 42.87584
	// fmt.Printf("x = %v/%T, y = %v/%T",x,x,y,y)
	x = 42
	y = 42.87584
	// x = 42.5344  wont work because of type
	fmt.Printf("x = %v/%T, y = %v/%T, z= %v/%T",x,x,y,y,z,z)
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}

// byte = uint8
// rune = int32
// runtime package, runtime.GOOS, runtime.GOARCH