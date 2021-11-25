package main

import (
	"fmt"
)

func main()  {
	// for with a single condition, like a while

	x:=1
	for x<10{
		fmt.Println(x)
		x++
	}

	// with for clause
	for i:=0; i<100; i++{
		fmt.Println(i)
	}

	j := 1
	// infinite loop
	for {
		if j<10{
			fmt.Println(j)
		}else {
			break
		}
		j++
	}
}