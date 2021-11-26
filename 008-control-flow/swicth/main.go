package main

import (
	"fmt"
)

func main()  {
	switch {
	case false:
		fmt.Println("This should not print")
	case (2==4):
		fmt.Println("This should not print2")
	case 3==3:
		fmt.Println("prints")
		fallthrough
	case 4==4:
		fmt.Println("also true, does it print?")
		fallthrough
	case 11==4:
		fmt.Println("Fallthrough, still prints")
	default:
		fmt.Println("default")
	}

	switch "Bond1"{
	case "Bond":
		fmt.Println("Bond")
	case "James":
		fmt.Println("James")
	default:
		fmt.Println("default")
	}

	switch '-'{
	case '1','-','=':
		fmt.Println("Yay")
	}
}