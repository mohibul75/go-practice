package otherpackages

import (
	"fmt"
	"strings"
)

func multiply(num1, num2 int) (int,int){
	return num1*num2, num1/num2
}

func functionWithUnusedVariable(num1, num2 int)int{
	div, _ := multiply(num1,num2)

	return div
}

func vardicFuntion(elements ...string) string{
	return strings.Join(elements,"")
}

func swap(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func add(a, b int) {
	fmt.Printf("the addition of two integer is : %d\n", a+b)
}

func findTheLargest(x, y, z int) int{

	var largest int

	if x>y && x>z {
		largest = x
	} else if y>x && y>z {
		largest = y
	} else{
		largest = z
	}

	return largest

}