package main

import(
	"fmt"
)

type author struct{
	name string
	branch string
	age int
	salary int
}

func (a *author) changeName(name string){
	(*a).name = name
}

func (a author) printAuth() {

	fmt.Printf("Author Name : %s\n", a.name)
	fmt.Printf("Author Branch : %s\n", a.branch)
	fmt.Printf("Author age : %d\n", a.age)
	fmt.Printf("Author salary : %d\n", a.salary)

}

func main(){

	auth:= author{
		name: "konal",
		branch: "CSE",
		age: 34,
		salary: 45678,
	}

	auth.printAuth()

	auth.changeName("Jamal")
	auth.printAuth()
}