package main

import "fmt"

func main(){
	a,b := 21,10

	if(a == b){
		fmt.Printf("first line - a == b \n")
	}else{
		fmt.Printf("first line - a != b \n")
	}

	
        if(a < b){
                fmt.Printf("second line - a < b \n")
        }else{
                fmt.Printf("second line - a >= b \n")
        }

	
        if(a > b){
                fmt.Printf("third line - a > b \n")
        }else{
                fmt.Printf("third line - a <= b \n")
        }
	
	a = 5
   	b = 20
	if(a <= b){
                fmt.Printf("fourth line - a <= b \n")
        }
	if(b >= a){
                fmt.Printf("fourth line - b >= a \n")
        }   	
}
