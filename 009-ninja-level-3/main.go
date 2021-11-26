package main

import (
	"fmt"
)
func main()  {
	// exercise1()
	// exercise2()
	// exercise3()
	// exercise4()
	// exercise5()
	// exercise6()
	// exercise7()
	// exercise8()
	// exercise9()
	exercise10()
}

func exercise1()  {
	for i:=0;i<=10000;i++{
		fmt.Println(i)
	}
}

func exercise2()  {
	for i:=65;i<=90;i++{
		fmt.Println(i)
		for j:=0;j<3;j++{
			fmt.Printf("\t%#U\n",i)
		}
	}
}

func exercise3()  {
	i := 1995
	for i < 2022{
		fmt.Println(i)
		i++
	}
}

func exercise4()  {
	i := 1995

	for{
		if i>2021{
			break
		}
		fmt.Println(i)
		i++
	}
}

func exercise5()  {
	for i:=10;i<101;i++{
		fmt.Println(i%4)
	}
}

func exercise6()  {
	if 10==10{
		fmt.Println("Match found")
	}
}

func exercise7()  {
	x := 11
	if x == 11{
		fmt.Println("X is 11")
	}else if x == 10{
		fmt.Println("X is 10")
	}else{
		fmt.Println("X is 9")
	}
}

func exercise8()  {
	switch {
	case 1==1:
		fmt.Println("1==1")
	case 1 != 1:
		fmt.Println("1 != 1")
	
	}
}

func exercise9()  {
	favSport := "swimming"
	switch favSport {
	case "skiing":
		fmt.Println("Go to mountains")
	case "swimming":
		fmt.Println("GO to pool")
	default:
		fmt.Println("What would you like to do?")
	}
}

func exercise10()  {
	
}