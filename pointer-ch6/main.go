package main

import (
	"fmt"
)

type person struct{
	firstName string
	middleName *string
	LastName string
}

func returnThePointer(middle string) *string{
	return &middle
}

func func1(g *int){
	x:= 10
	g= &x
}

func main(){

	x:=10
	pointerX:= &x

	fmt.Printf("%d\n", pointerX)
	fmt.Printf("%d\n", *pointerX)

	var y = new(int)
	*y= 6
	y= &x

	fmt.Println(y)
	fmt.Println(*y)

	p:= person{
		firstName: "khan",
		middleName: returnThePointer("hjkl"),
		LastName: "tyu",
	}

	fmt.Println(p)

	var g *int
	func1(g)
	fmt.Println(g)

}