package main

import "fmt"

func main(){
	a,b,c := 21,10,0

	c = a + b
	fmt.Printf("first line,c is %d\n",c)

	c = a - b
	fmt.Printf("second line,c is %d\n",c)

	c = a * b
	fmt.Printf("third line,c is %d\n",c)

	c = a / b
	fmt.Printf("forth line,c is %d\n",c)

	c = a % b
	fmt.Printf("fifth line,c is %d\n",c)

	a ++
	fmt.Printf("sixth line,a is %d\n",a)
	
	a = 21
	a --
	fmt.Printf("seventh line,a is %d\n",a)
}
