package main

import (
	"fmt"
	"time"
)

type Counter struct{
	total int
	lastUpdated time.Time
}

func doUpdateWrong(c Counter){
	c.Increment()
	fmt.Println("in doupdateWrong :", c.toString())
}

func doUpdateRight(c *Counter){
	c.Increment()
	fmt.Println("in doUpdateRight :", c.toString())
}

func (c *Counter) Increment(){
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) toString() string{
	return fmt.Sprintf("total : %d, last update: %v", c.total, c.lastUpdated)
}

func main(){
	var c Counter
	doUpdateWrong(c)
	fmt.Println("in main : ",c.toString())
	doUpdateRight(&c)
	fmt.Println("in main for 2nd",c.toString())
}