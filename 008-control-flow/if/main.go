package main

import (
	"fmt"
)

func main()  {
	// conditional
	if true{
		fmt.Println("001")
	}
	if false{
		fmt.Println("002")
	}
	if !true{
		fmt.Println("003")
	}
	if !false{
		fmt.Println("004")
	}
	if 2==2{
		fmt.Println("005")
	}
	if 2!=2{
		fmt.Println("006")
	}

	//initialization
	if i:=42;i==42{
		fmt.Println("42==42")
	}

	// if-else-if-else
	x := 42
	if x == 40{
		fmt.Println("Our value is 40")
	}else if x == 41{
		fmt.Println("Our value is 41")
	}else{
		fmt.Println("Our value is 42")
	}

}