package main

import (
	"fmt"
)

type person struct{
	first string
	last string
	icecreams []string
}

type vehicle struct{
	doors int
	color string
}

type truck struct{
	vehicle
	fourWheel bool
}

type sedan struct{
	vehicle
	luxury bool
}

func main()  {
	// exercise1()
	// exercise2()
	// exercise3()
	exercise4()
}

func exercise1()  {
	p1 := person{
		first: "Mansi",
		last: "Panpaliya",
		icecreams: []string{"chocolate", "dryfruit", "Kesar Pista"},
	}

	p2 := person{
		first: "Shlok",
		last: "Panpaliya",
		icecreams: []string{"Mango", "Raspberry"},
	}

	fmt.Println(p1.first, p1.last)
	fmt.Println("\tIcecreams: ")
	for _,v := range p1.icecreams{
		fmt.Println("\t",v)
	}

	fmt.Println(p2.first, p2.last)
	fmt.Println("\tIcecreams: ")
	for _,v := range p2.icecreams{
		fmt.Println("\t",v)
	}
}

func exercise2()  {
	m := map[string]person{}
	m["Panpaliya"] = person{
		first: "Mansi",
		last: "Panpaliya",
		icecreams: []string{"chocolate", "dryfruit", "Kesar Pista"},
	}

	for _,v := range m{
		fmt.Println(v.first, v.last)
		fmt.Println("Icecreams:")
		for _,f := range v.icecreams{
			fmt.Println(f)
		}
	}

	p1 := person{
		first: "Mansi",
		last: "Panpaliya",
		icecreams: []string{"chocolate", "dryfruit", "Kesar Pista"},
	}

	p2 := person{
		first: "Shlok",
		last: "Manish",
		icecreams: []string{"Mango", "Raspberry"},
	}

	m1 := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _,v := range m1{
		fmt.Println(v.first, v.last)
		fmt.Println("Icecreams:")
		for _,f := range v.icecreams{
			fmt.Println(f)
		}
	}

}

func exercise3()  {

	v1 := vehicle{
		doors: 4,
		color: "black",
	}

	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	} 

	fmt.Println(v1, v1.doors, v1.color)
	fmt.Println(t1, t1.doors, t1.color, t1.fourWheel)
	fmt.Println(s1, s1.doors, s1.color, s1.luxury)
}

func exercise4()  {
	t := struct{
		carTypes []string
		car map[string]string
	}{
		carTypes: []string{"hatchback","sedan","SUV"},
		car: map[string]string{"Brezza":"white", "Dzire":"Pearl","800":"white"},
	}

	fmt.Println(t,t.carTypes,t.car)
}