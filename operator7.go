package main

import "fmt"

func main(){
	var a int = 20
	var b int = 10
	var c int = 15
	var d int = 5
	var e int

	e = (a + b) * c / d	
	fmt.Printf("(a + b) * c /d value is %d\n",e)

	e = ((a + b) * c) / d
	fmt.Printf("e = ((a + b) * c) / d value is %d\n",e)
	
	e = (a + b) * (c / d)
	fmt.Printf("e = (a + b) * (c / d) value is %d\n",e)
	
	e = a + (b * c) / d
	fmt.Printf("e = a + (b * c) / d value is %d\n",e)
}
