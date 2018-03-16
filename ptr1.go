package main

import "fmt"

func main(){

	var a int = 10
	var p *int
	p = &a

	fmt.Printf("this var address :%x\n",&a)
	fmt.Printf("p  is :%x\n",p)
	fmt.Printf("*p is :%d\n",*p)

	var ptr *int
	fmt.Printf("ptr  is :%x\n",ptr)
}
