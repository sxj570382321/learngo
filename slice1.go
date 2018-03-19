package main

import "fmt"

func main(){

	var slice1 = [4] int{1,2,3,4}
	fmt.Println(slice1)

	slice2 := [] int{1,2,3,4,5}
	fmt.Println(slice2)

	var slice3[] int = make([]int,3) // <==> var slice3 = make([]int,3)
	fmt.Println(slice3)

	slice4 := []int {1,2,3,4,5,6}
	fmt.Println(slice4)

	arr := [] int {1,2,3,4,5,6,7,8,9,10}
	slice5 := arr[2:5]  //from 2 to 5,include 2,but not include 5
	fmt.Println(slice5)

	slice6 := make([]int,4,10)
	fmt.Println(slice6)
}
