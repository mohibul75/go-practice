package printpackage

import (
	"fmt"
)

func printResult(res int){
	fmt.Printf("the output is %d\n", res)
}

func printDate(day int){

	switch day{
	case 1:
		fmt.Printf("monday\n")
	case 2:
		fmt.Printf("Tuesday\n")
	case 3:
		fmt.Printf("Wednesday\n")
	case 4:
		fmt.Printf(("Thursday\n"))
	case 5:
		fmt.Printf("Friday\n")
	case 6:
		fmt.Printf("Staturday\n")
	case 7:
		fmt.Printf("Sunday\n")
	}

}