package main

import "fmt"

func main(){
	var a uint = 60
	var b uint = 13
	var c uint = 0

	c = a & b
	fmt.Printf("first line,c is %d\n",c)

	c = a | b
        fmt.Printf("second line,c is %d\n",c)

	c = a ^ b
        fmt.Printf("third line,c is %d\n",c)

	c = a << 2
        fmt.Printf("fourth line,c is %d\n",c)

	c = a >> 2
        fmt.Printf("fifth line,c is %d\n",c)
}
