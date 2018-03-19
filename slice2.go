package main

import "fmt"

func main(){

	var numbers = make([]int,3,10)
	printSlice(numbers)
	
	var numbers1[] int
        printSlice(numbers1)
	if(numbers1 == nil){
		fmt.Printf("this slice is empty\n")
	}
}


func printSlice(x[] int){
	fmt.Printf("len=%d,cap=%d, slice=%v\n",len(x),cap(x),x)
}
