package main

import "fmt"

func main(){
	var a int = 4
	var b int32
	var c float32
	var ptr* int

	ptr = &a
	fmt.Printf("a is %d,and ptr is %x\n",a,ptr)
	fmt.Printf("*ptr is %d\n",*ptr)
	
	fmt.Printf("first line -a type is = %T\n",a)
	fmt.Printf("second line -b type is = %T\n",b)
	fmt.Printf("third line -c type is = %T\n",c)
}
