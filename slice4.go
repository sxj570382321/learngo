package main

import "fmt"

func main(){

        var numbers[] int 
        printSlice(numbers)

	numbers = append(numbers,0)
	printSlice(numbers)
	
	numbers = append(numbers,1)
        printSlice(numbers)

	numbers = append(numbers,2,3,4)
        printSlice(numbers)

	var numbers2 = append(numbers,5,6,7,8)
        printSlice(numbers2)

        numbers1 := make([]int,len(numbers),(cap(numbers)) * 2)
        printSlice(numbers1)

	copy(numbers1,numbers)
	printSlice(numbers1)

}


func printSlice(x[] int){
        fmt.Printf("len=%d,cap=%d, slice=%v\n",len(x),cap(x),x)
}

