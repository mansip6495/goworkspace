package main

import (
	"fmt"
)

// untyped constants
const a = 42
const b = 42.78
const c = "James Bond"

// Types constants
// const(
// 	a int = 42
// 	b floadt64 = 42.78
// 	c string = "James Bond" 
// )

func main()  {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",c)
}