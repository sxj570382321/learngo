package main

import "fmt"

func main() {

	var sum int = 17
	var count int = 5
	var mean float32
	var a = 65

	mean = float32(sum) / float32(count)
	fmt.Printf("mean is %f\n", mean)
	b := string(a)
	fmt.Println("b is:", b)
}
