package main


import (
	"fmt"
)

const(
	a = 42
	b = "James Bond"
	c int = 43
)

const(
	_ = iota
	y1 = 2017 + iota
	y2 = 2017 + iota
	y3 = 2017 + iota
	y4 = 2017 + iota
)

func main()  {
	// exercise1()
	// exercise2()
	// exercise3()
	// exercise4()
	// exercise5()
	// exercise6()
}

func exercise1(){
	x := 42
	fmt.Printf("\n %d \t %b \t %x", x, x, x)
}

func exercise2(){
	a := 7
	b := 42

	c := a == b
	d := a <= b
	e := a >= b
	f := a != b
	g := a < b
	h := a > b

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}

func exercise3()  {
	fmt.Println(a,b,c)
}

func exercise4()  {
	x := 32
	fmt.Printf("\n %d \t %b \t %x", x, x, x)

	y := x << 1
	fmt.Printf("\n %d \t %b \t %x", y, y, y)
}

func exercise5()  {
	s := `Hello
world!`
	fmt.Println(s)
}

func exercise6()  {
	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
	fmt.Println(y4)
}
