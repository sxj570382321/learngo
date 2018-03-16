package main

import "fmt"

const MAX int = 3
func main(){
	var a = [] int{10,100,200}
//	a := [] int{10,100,200}
	var i int

	for i = 0;i <MAX;i ++{
		fmt.Printf("a[%d] = %d\n",i,a[i])
	}


	fmt.Println("#####################")
	var ii int
	var ptr [MAX] *int
	for ii = 0; ii < MAX; ii ++{
		ptr[ii] = &a[ii]
	}

	for ii = 0; ii < MAX; ii ++{
		fmt.Printf("a[%d] = %d\n",ii,*ptr[ii])
	}
}
