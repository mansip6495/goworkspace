package main

import (
	"fmt"
)

type person struct{
	first string
	last string
}

type secretAgent struct{
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code }

func (s secretAgent) speak()  {
	fmt.Println("SecretAgent - I am,", s.first, s.last)
}

func (p person) speak()  {
	fmt.Println("Person - I am,", p.first, p.last)
}

// Interface allows to define behavior, also allows to do polymorphism
// anything that has method speak is also of type human
// a value can be of multiple type
type human interface{
	speak()
}

func bar(h human) {
	switch h.(type){
	case person:
		fmt.Println("I called human", h.(person).first, h.(person).last)
	case secretAgent:
		fmt.Println("I called human", h.(secretAgent).first, h.(secretAgent).last)
	}
	// h.(secretAgent).first - is called assertion
	// fmt.Println("I called human:",h)
}

func main()  {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}

	p1 := person{
		first: "Miss",
		last: "MoneyPenny",
	}

	fmt.Printf("%T\n",sa1)

	//polymorphism
	// As inteface human has function speak, any value that has method speak
	// will aslo be of type human
	bar(sa1)
	bar(p1)
}