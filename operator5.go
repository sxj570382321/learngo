package main

import "fmt"

func main(){
	var a int = 21
	var c int

	c = a
	fmt.Printf("first line - operator =,c is %d\n",c)

	c += a
	fmt.Printf("second line - operator +=,c is %d\n",c)

	c -= a
	fmt.Printf("third line - operator -=,c is %d\n",c)

	c *= a
	fmt.Printf("foruth line - operator *=,c is %d\n",c)


	c /= a
	fmt.Printf("fifth line - operator /=,c is %d\n",c)

	
	c = 200
	c <<= 2
	fmt.Printf("sixth line - operator <<=,c is %d\n",c)
	
	c >>= 2
	fmt.Printf("seventh line - operator >>=,c is %d\n",c)


	
        c &= 2
        fmt.Printf("eighth line - operator >&=,c is %d\n",c)

	
        c ^= 2
        fmt.Printf("nineth line - operator ^=,c is %d\n",c)

	
        c |= 2
        fmt.Printf("tenth line - operator |=,c is %d\n",c)



}
