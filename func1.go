package main

import (
	"fmt"
	"math"
)

func max(num1 int,num2 int) int{
	var result int
	
	if(num1 > num2){
		result = num1
	}else{
		result = num2
	}	
	return result
}


func main(){
	var result int
	result = max(10,9)
	fmt.Println("max 10,9 is:",result)

	x,y:= swap("10","9")
	fmt.Printf("swap 10,9 is :%s,%s\n",x,y)

	getSquareRoot:=func(x float64) float64{  //go,func as return value
			return math.Sqrt(x)
			}	
	fmt.Println(getSquareRoot(9))

	nextNumber := getSequence()

	fmt.Println(nextNumber())

	fmt.Println(nextNumber())

	fmt.Println(nextNumber())

	fmt.Println("###########")
	
	fmt.Println(nextNumber())

	fmt.Println(nextNumber())

	fmt.Println("$$$$$$$$$$$$$")

	nextNumber1 := getSequence()

	fmt.Println(nextNumber1())

	fmt.Println(nextNumber1())

	
	addfun := add(1,2)
	fmt.Println(addfun())
	
	fmt.Println(addfun())

	fmt.Println(addfun())
}


func swap(x,y string)(string,string){
	return y,x
}


func getSequence() func()int{ //niminghanshu,not have function name,but func is need

	i := 0
	return func()int{
	i += 1
	return i
	}

}


func add(x,y int) func()(int,int){ //niminghanshu,not have function name,but func is need
	i := 0

	return func()(int,int){
	i ++
	return i + x,i + y
	}
}
