package main

import "fmt"

type shape interface{
	area() int
}

type rectangle struct{
	length int
	width int
}

type circle struct{
	radius int
}

type saqure struct{
	length int
}

func (s *rectangle) area() int{
	return s.length*s.width
}

func (s *circle) area() int{
	return int(2*3.1416*float64(s.radius))
}

func(s *saqure) area() int{
	return s.length*s.length
}

func getArea(s shape) int{
	return s.area()
}

func main(){
	r:= rectangle{20,25}
	c:= circle{10}
	s:= saqure{5}

	shapes:=[]shape{&r, &c, &s}

	for _, sh:= range shapes{
		fmt.Println("Area", getArea(sh))
	}

}