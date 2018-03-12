package main

import "fmt"

const (
      Unknow = 0
      Female = 1
      Male = 2
)

func main(){

	const LENGTH =10
	const WIDTH = 5
	var area int
	const a,b,c = 1,false,"str"

	area = LENGTH * WIDTH
	fmt.Printf("AREA is :%d\n",area)
	fmt.Println()
	fmt.Println(a,b,c)
	println(area)
}
