package main

import (
	"fmt"
	"reflect"
)

type User struct {
	id int32
}

func main() {
	var num float64 = 1.2345
	fmt.Println("old value of pointer:", num)

	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()

	fmt.Println("type of pointer:", newValue.Type())
	fmt.Println("settability of pointer:", newValue.CanSet())

	newValue.SetFloat(77)
	fmt.Println("newValue of pointer:", num)

	var us User = User{1}
	pointer1 := reflect.ValueOf(&us)
	newValue1 := pointer1.Elem()
	fmt.Println(newValue1)
	fmt.Println("settability of pointer:", newValue1.CanSet())

	//	newValue1.SetInt(2)
	fmt.Println("newValue of pointer:", us)
}
