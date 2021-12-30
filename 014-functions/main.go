package main

import (
	"fmt"
)

func main()  {
	foo()
	bar("James")
	s1 := woo("Money Penny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x,y)

	// Variadic parameter
	variadic1(2,3,4,5,6,7,8,9)

	// Unfurling a slice
	xi := []int{2,3,4,5,6,7,8,9}
	unfurlSlice(xi...)

	//Defer
	defer defer1()
	defer2()

	// Anonymous func
	func ()  {
		fmt.Println("Anonymous function")
	}()

	func (x int)  {
		fmt.Println("Anonymous func with parameters:", x)
	}(42)

	// Func experssion
	// just like int can be assigned to a variable, so is func
	f := func ()  {
		fmt.Println("my first func expression")
	}
	f()

	// return func from func
	s := retunFunc()
	sa := s()
	fmt.Printf("%T\n",s)
	fmt.Println(sa)

	// callback: pass a func as argument
	i := []int{1,2,3,4,5,6,7,8,9,}
	fmt.Println("Sum of numbers from 1-9",sum1(i ...))
	fmt.Println("Sum of even numbers from 1-9", even(sum1, i ...))

	// closure - scope limit
	{
		y := 56
		fmt.Println("Closure",y)
	}

}

// func (r receiver) identifier (parameters) (returns) { code }
// Parameters: we define our func with parameters
// Arguments: we call our func and pass in arguments
func foo() {
	fmt.Println("hello from foo")
}

// Everything in GO is PASS BY VALUE
func bar(s string)  {
	fmt.Println("Hello, ",s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo ",s)
}

func mouse(first string, last string) (string,bool) {
	a := fmt.Sprint(first, last, `, says "Hello"`)
	return a, false
}

// the variadic parameter should be the final parameter
func variadic1(x ...int)  {
	fmt.Printf("%T: %v\n", x,x)
	sum := 0
	for _,v := range x{
		sum += v
	}
	fmt.Println("From variadic1",sum)
}

// Unfurling a slice
func unfurlSlice(x ...int)  {
	// x := []int{2,3,4,5,6,7,8,9}
	fmt.Printf("%T: %v\n", x,x)
	sum := 0
	for _,v := range x{
		sum += v
	}
	fmt.Println("From unfurlSlice",sum)
}

//defer
//ex when you open a file, write the close file with a defer
// as its deferred it will execute at the end of function before exiting
func defer1(){
	fmt.Println("Defer1")
}

func defer2() {
	fmt.Println("Defer2")
}

// return func from func
func retunFunc() func() int {
	return func () int  {
		return 42
	}
}

// callback: pass a func as parameter
func sum1(x ...int) int {
	total := 0
	for _,v := range x{
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _,v := range vi{
		if v % 2 == 0{
			yi = append(yi, v)
		}
	}
	return f(yi ...)
}

// Recursion: func calls itself
func factorial(n int) int {
	if n == 0{
		return 1
	}
	return n * factorial(n-1)
}