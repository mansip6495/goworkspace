package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, I am at this amazing go learning class of Todd McLeod")
	foo()
	fmt.Println("Something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I am in foo")
}

func bar()  {
	fmt.Println("and then we exited!")
}
// control flow:
// 1] Sequence:
// 2] Loop:
// 3] Conditional
