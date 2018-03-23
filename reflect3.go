package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name:", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs")
}

func main() {
	user := User{1, "sxj", 30}

	getValue := reflect.ValueOf(user)

	methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
	args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(20)}
	methodValue.Call(args)

	methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
	args = make([]reflect.Value, 0)
	methodValue.Call(args)

}
