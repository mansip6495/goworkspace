package main

import(
	"fmt"
)

type person struct{
	first string
	last string
	age int
}

func (p person) speak()  {
	fmt.Println("Hi, this is ", p.first, p.last, "and I am ", p.age, "years old")
}

type square struct{
	length int
}

type circle struct{
	radius int
}

func (s square) area()  {
	fmt.Println("Square Area: ", s.length*s.length)
}

func (c circle) area()  {
	fmt.Println("Circle Area: ", 22 * c.radius * c.radius / 7)
}

type shape interface{
	area()
}

func main()  {
	// Exercise 1
	fmt.Println(foo1())
	fmt.Println(bar1())

	// Exercise 2
	x2 := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(foo2(x2 ...))
	fmt.Println(bar2(x2))

	// Exercise 3
	defer foo3()

	// Exercise 4
	p4 := person{
		first: "James",
		last: "Bond",
		age: 32,
	}
	p4.speak()

	// Exercise 5
	s5 := square{
		length: 5,
	}

	c5 := circle{
		radius: 7,
	}
	
	info5(s5)
	info5(c5)

	// Exercise 6
	func ()  {
		fmt.Println("Anonymous function")
	}()
	
	// Exercise 7
	x7 := func ()  {
		fmt.Println("Assign function to variable")
	}
	x7()

	// Exercise 8
	x8 := foo8()
	fmt.Println(x8())

	// Exercise 9
	x9 := []int{1,2,3,4,5,6,7,8,9,}
	fmt.Println(sum9(x9 ...))
	fmt.Println(even(sum9, x9 ...))
}

func foo1()int{
	return 1
}

func bar1() (int,string) {
	return 2, "Hello, World!"
}

func foo2(x ...int)int{
	total := 0
	for _,v := range x{
		total += v
	}
	return total
}

func bar2(x []int) int {
	total := 0
	for _,v := range x{
		total += v
	}
	return total
}

func foo3()  {
	fmt.Println("FOO ran: defer demo")
}

func info5(s shape){
	s.area()
}

func foo8() func() int {
	return func () int  {
		fmt.Println("Returning a func")
		return 8
	}
}

func sum9(x ...int) int {
	total := 0
	for _,v := range x{
		total += v
	}
	return total
}

func even(f func(t ...int) int, v ...int) int  {
	var x []int
	for _,v1 := range v{
		if v1 % 2 == 0{
			x = append(x, v1)
		}
	}
	return f(x ...)
}





