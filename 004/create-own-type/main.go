package main
import "fmt"

var a int
type hotdog int
var b hotdog

func main()  {
	a = 42
	b= 43
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	fmt.Printf("b=%v is of type %T \n",b,b)
	//a = b wont work as a is of type int and cannot be assigned a value of type hotdog
	a = int(b)
	// CONVERT a variable of type hotdog to type int
	fmt.Println(a)
}