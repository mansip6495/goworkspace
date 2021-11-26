package main

import (
	"fmt"
)

func main()  {
	for i:=33; i<=122; i++{
		fmt.Printf("\n%d -> %c \t %#U",i,i,i)
	}
}