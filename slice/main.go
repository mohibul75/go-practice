package main

import (
	"fmt"
	"go-packages/otherPackages"
	"go-packages/printPage"
)

func main(){
	arr:=[4]int {2,3,4,5}

	var i int

	for i=0; i<4; i++ {
		fmt.Println(arr[i])
	}

	var a int = 89
	var b int = 567
	add(a,b)



	var str string = "dhhjjk"
	fmt.Printf("the value of string variable : %s\n", str)

	bl:= true
	fmt.Printf("the value of boolean variable : %t\n", bl)

	x:= 5
	y:= 6
	z:= 10
	fmt.Printf("the largest number is %d\n", findTheLargest(x,y,z))
	

	strArr := [5]string{"a", "b", "c", "d", "e"}
	for m,n:= range strArr {
		fmt.Printf("%d is for %s\n",m,n)
	}

	day:= 4
	printDate(day)

	temp1:= 4
	temp2:= 8
	fmt.Printf("Before swapping %d and %d\n", temp1, temp2)
	swap(&temp1, &temp2)
	fmt.Printf("After swapping %d and %d\n", temp1, temp2)

	str2:= vardicFuntion("a","b", "c")
	fmt.Printf("the joining string is %s\n", str2)

	 mul:=  functionWithUnusedVariable(4,5)
	 printResult(mul)
	
}