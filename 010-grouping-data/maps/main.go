package main

import (
	"fmt"
)

func main()  {
	//map: key value pairs, unordered list
	m := map[string]int{
		"James": 32,
		"Miss Moneypenny": 27,
	}    // composite literal
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])  // If key doesnt exist, then it returns 0

	// check if a key exists with comma ok idiom
	v,ok := m["Barnabas"]
	fmt.Println(v,ok)

	if v,ok := m["Miss Moneypenny"]; ok{
		fmt.Println("This is the if print:",v)
	}

	// Add elements to map
	m["Barnabas"] = 25

	for k,v := range(m){
		fmt.Println(k, "-> ",v)
	}

	// Delete an entry
	delete(m,"Barnabas")
	fmt.Println(m)
}