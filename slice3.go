package main

import "fmt"

func main(){

       	numbers := []int {1,2,3,4,5,6,7,8,9}
	printSlice(numbers)
	fmt.Println("numbers ==",numbers)
	fmt.Println("numbers[1:4] ==",numbers[1:4])
	fmt.Println("numbers[:3] ==",numbers[:3])
	fmt.Println("numbers[4:] ==",numbers[4:])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	numbers2 := numbers[:2]
	printSlice(numbers2) //len 2,but cap 9

	numbers3 := numbers[2:5]
        printSlice(numbers3) //len 3,but cap 7


}


func printSlice(x[] int){
        fmt.Printf("len=%d,cap=%d, slice=%v\n",len(x),cap(x),x)
}
