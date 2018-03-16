package main

import "fmt"

func main(){
	var a int = 100
	var b int = 200

	fmt.Printf("before swap,a is %d\n",a)
	fmt.Printf("before swap,b is %d\n",b)

	swap(&a,&b)

	
        fmt.Printf("after swap,a is %d\n",a)
        fmt.Printf("after swap,b is %d\n",b)

}

func swap(x,y *int) {

	var temp int
	temp = *x
	*x = *y
	*y = temp
}
