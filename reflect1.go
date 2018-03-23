package main

import (
	"fmt"
	"reflect"
)

type student struct {
	id   int32
	name string
	age  rune
}

func (su student) hello() {
	fmt.Println("hello student")
}

func (su student) hello1() {
	fmt.Println("hello1 student")
}

func main() {
	/*
		var a float32 = 1.234000
		fmt.Printf("type:%s\n", reflect.TypeOf(a))
		fmt.Printf("value:%v\n", reflect.ValueOf(a)) //%f print 1.234000,but %v print 1.234
		fmt.Println("hello world float32!")

		var stu student = student{1, "wanghui", 20.000}
		fmt.Printf("type:%s\n", reflect.TypeOf(stu))
		fmt.Printf("vlaue:%v\n", reflect.ValueOf(stu))
		fmt.Println("hello world student!")

		var num float64 = 1.2345
		pointer := reflect.ValueOf(&num)
		value := reflect.ValueOf(num)

		fmt.Println("pointer:", pointer)
		fmt.Println("value:", value)

		convertPOinter := pointer.Interface().(*float64)
		converValue := value.Interface().(float64)
		fmt.Println("convertPOinter:", convertPOinter)
		fmt.Println("convertValue:", converValue)

		var stu1 student = student{2, "zhangsan", 21}
		point_stu1 := reflect.ValueOf(&stu1)
		value_stu1 := reflect.ValueOf(stu1)
		fmt.Println("point_stu1:", point_stu1)
		fmt.Println("value_stu1:", value_stu1)

		convertPOinter_stu1 := point_stu1.Interface().(*student)
		converValue_stu1 := value_stu1.Interface().(student)
		fmt.Println("convertPOinter_stu1:", convertPOinter_stu1)
		fmt.Println("converValue_stu1:", converValue_stu1)
	*/

	var stu1 student = student{2, "zhangsan", 21}
	//	DoFiledAndMethod(stu1)
	DoFiledAndMethod1(stu1)
}

func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("get tyepe is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i)
		fmt.Printf("field,name:%s,type:%v,value:%v\n", field.Name, field.Type, value)
	}

	fmt.Println("2222222222")
	i := getType.NumMethod()
	fmt.Println(i)
	for i = 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Println("1111111111111")
		fmt.Printf("method,name:%s,type:%v\n", method.Name, method.Type)
	}
}

func DoFiledAndMethod1(input interface{}) {
	getVaue := reflect.ValueOf(input)

	for i := 0; i < getVaue.NumField(); i++ {
		value := getVaue.Field(i)
		fmt.Println(value)
	}

	fmt.Println("33333333333")
	i := getVaue.NumMethod()
	fmt.Println(i)
}
