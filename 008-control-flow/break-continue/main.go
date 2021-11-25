package main

import "fmt"

func main(){
	x := 41 / 40
	y := 43 % 40
	fmt.Println(x,y)

	i := 0
	for{
		i++
		if i>100{
			break
		}
		if i%2!=0{
			continue
		}
		fmt.Println(i)
		
	}
}