package main

import(
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
	a := [5]int{1,2,3,4,5}
	for i,v := range a{
		fmt.Println(i, "->", v)
	}
	fmt.Printf("%T",a)
}

func exercise2()  {
	a := []int{1,2,3,4,5,6,7,8,9,10}
	for i,v := range a{
		fmt.Println(i, "->", v)
	}
	fmt.Printf("%T",a)
}

func exercise3()  {
	a := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(a[:5])
	fmt.Println(a[5:])
	fmt.Println(a[2:7])
	fmt.Println(a[1:6])
}

func exercise4()  {
	a := []int{42,43,44,45,46,47,48,49,50,51}
	a = append(a,52)
	fmt.Println(a)
	a = append(a,53,54,55)
	fmt.Println(a)
	x := []int{56,57,58,59,60}
	a= append(a,x...)
	fmt.Println(a)
}

func exercise5()  {
	a := []int{42,43,44,45,46,47,48,49,50,51}
	a = append(a[:3],a[6:]...)
	fmt.Println(a)
}

func exercise6()  {
	a := make([]string,4,4)
	state := []string{"AZ","CA","FL","NY"}
	for i,v := range state{
		a[i] = v
	}
	fmt.Println(a,len(a),cap(a))
}

func exercise7()  {
	a := []string{"James","Bond","Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Hellooo, James"}

	result := [][]string{a,b}

	for i,v := range result{
		for j,v1 := range v{
			fmt.Println(i,j,v1)
		}
	}
}

func exercise8()  {
	m := map[string][]string{}
	m["bond_james"] = []string{"Shaken, not stirred","Martinis","Women"}
	m["moneypenny_miss"] = []string{"james bond","Literature","Computer Science"}
	m["no_dr"] = []string{"Being Evil","Ice Cream", "Sunsets"}

	for k,v := range m{
		for _,v1 := range v{
			fmt.Println(k, v, v1)
		}
	}
}

func exercise9()  {
	m := map[string][]string{}
	m["bond_james"] = []string{"Shaken, not stirred","Martinis","Women"}
	m["moneypenny_miss"] = []string{"james bond","Literature","Computer Science"}
	m["no_dr"] = []string{"Being Evil","Ice Cream", "Sunsets"}
	m["fleming_in"] = []string{"steak","cigar","espionage"}
	fmt.Println(m)
}

func exercise10()  {
	m := map[string][]string{}
	m["bond_james"] = []string{"Shaken, not stirred","Martinis","Women"}
	m["moneypenny_miss"] = []string{"james bond","Literature","Computer Science"}
	m["no_dr"] = []string{"Being Evil","Ice Cream", "Sunsets"}
	m["fleming_in"] = []string{"steak","cigar","espionage"}
	delete(m,"fleming_in")
	fmt.Println(m)
}