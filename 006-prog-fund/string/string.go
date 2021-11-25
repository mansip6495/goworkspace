package main

import (
	"fmt"
)

func main()  {
	// strings can be created with "" or ``. `` can take \n as well
	s := "Hello World"
	fmt.Println(s)
	bs := []byte(s)
	fmt.Printf("%v, %T", bs, bs)

	for i:=0; i< len(s); i++{
		fmt.Printf("%#U ", s[i])
	}
}
// https://go.dev/blog/strings