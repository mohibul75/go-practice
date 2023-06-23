package main

import (
	"fmt"
)

func main(){
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2] //1,2
	z := x[2:] //3,4
	fmt.Println(cap(x), cap(y), cap(z)) //5,5,5
	y = append(y, 30, 40, 50) // 1,2,30,40,50
	x = append(x, 60) // 1,2,30,40,50
	z = append(z, 70) //30, 40, 70
	fmt.Println("x:", x)  //1,2,30,40,50.70
	fmt.Println("y:", y) // 1,2,30,40,70
	fmt.Println("z:", z) //30, 40, 70


	var nilMap map[string]int 
	fmt.Println(nilMap)

	// nilMap["a"]= 5 //does not work
	// nilMap["b"]= 89 // does not work

	fmt.Println(nilMap["a"])

	another_nilMap:= map[string]int{}
	fmt.Println(another_nilMap)

	temas:= map[string][]string{
		"a":[]string{"Fred", "fhbn", "fhj"},
		"b":[]string{"djnd","dnjd","iopp"},
		"c":[]string{"yuuk","iop","ujno"},
	}
	fmt.Println(temas)

	ages:= make(map[int][]string, 10)
	fmt.Println(ages)

	m:= map[string]int{
		"hello": 90,
		"b": 89,
	}
	v, ok:= m["hello"]
	fmt.Println("for hello",v,ok)
	v,ok= m["world"]
	fmt.Println("for world", v, ok)

	st:= map[int]bool{}
	vl:= []int{5,6,3,4,3,23,45,23,4}
	for _,v:=range vl{
		st[v]=true
	}

	fmt.Println(len(vl), len(st))
	fmt.Println("st[3]", st[3])
	fmt.Println("st[500]", st[500])

}