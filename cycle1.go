package main

import "fmt"
//import "math"

func main(){
	var C,c int
	C = 1

	A: for (C < 100){
		C++
		for c = 2; c < C;c ++	{
			if C % c == 0{
				goto A
			}
		}
		fmt.Println(C,"is sushu")
	}
}
