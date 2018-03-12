package main

import "fmt"

func main(){
	a,b := true,false
	if( a && b){
		fmt.Printf("first line,cond is true\n")
	}
	if( a || b){ 
                fmt.Printf("second line,cond is true\n")
        }

	a = false
	b = true
	if(a && b){
		fmt.Printf("third line,cond is true\n")
	}else{
		fmt.Printf("third line,cond is false\n")
	}

	if(!(a && b)){ 
                fmt.Printf("fourth line,cond is true\n")
        }
}
