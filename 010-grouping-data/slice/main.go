package main

import (
	"fmt"
)

func main()  {
	// x := type{values} //composite literal
	x := []int{4, 5, 7, 8, 42} 
	fmt.Println(x, len(x), cap(x))
	// a SLICE allows you to group together VALUES of the same TYPE

	// Loop over slice
	for i,v := range x{
		fmt.Println(i,v)
	}

	for i:=0;i<len(x);i++{
		fmt.Println(i, " -> ", x[i])
	}

	// Slicing a slice
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	// append to a slice
	x = append(x,77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 897}
	x = append(x,y...)
	fmt.Println(x)

	// delete from a slice
	// delete 7,8
	x = append(x[:2],x[4:]...)
	fmt.Println(x)

	// make
	// If you know the size of underlying array use make. It saves the time to create a new 
	// bigger array and delete the own one stuff like that. If you dont know the size of 
	// underlying array, u can use composite literal 
	a := make([]int, 10,100)
	fmt.Println(a,len(a),cap(a))
	a[0] = 42
	a[9] = 999
	fmt.Println(a,len(a),cap(a))
	// will error out -> a[10]=1000
	a = append(a,3999)
	fmt.Println(a,len(a),cap(a))

	// If ur cap is 12 and u add the the 13th element, then the new underlying array needs to 
	// be initialized and the old one is thrown away

	// Multi-dimensional slice
	jb := []string{"James", "Bond", "Choclate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb,mp}
	fmt.Println(xp)

	for i,v := range xp{
		fmt.Println(i,v)
		for j:=0; j<len(v);j++{
			fmt.Println(xp[i][j])
		}
		
	}
}
