package main

import (
	"fmt"
)

//DS allows to compose values of diff type, aggregate DS
type person struct{
	first string
	last string
}

// no need of an identifier for embedded struct - Anonymous field
type secretAgent struct{
	person  
	ltk bool
}

type secretAgentcollision struct{
	person  
	ltk bool
	first string
}

func main()  {
	p1 := person{
		"James",
		"Bond",
	}
	p2 := person{
		"Miss", 
		"Moneypenny",
	}
	fmt.Println(p1,p1.first,p1.last)
	fmt.Println(p2,p2.first,p2.last)

	// Embedded struct
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1, sa1.first, sa1.last, sa1.ltk)

	sa2 := secretAgentcollision{
		person: person{
			first: "James",
			last: "Bond",
		},
		first: "James1",
		ltk: true,
	}

	fmt.Println(sa2, sa2.first, sa2.last, sa2.ltk, sa2.person.first)

	//Anonymous Structs
	p3 := struct{
		first string
		last string
	}{
		first: "Mansi",
		last: "Panpaliya",
	}
	fmt.Println(p3)
}

