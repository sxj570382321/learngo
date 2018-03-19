package main

import "fmt"

type Phone interface{
	call()
}

type NokiaPhone struct{


}

func (nk NokiaPhone) add() int{
	fmt.Println("this is add of NokiaPhone")
	return 0
}


func (nokiaPhone NokiaPhone) call(){
	fmt.Println("I am Nokia,I can call you!")
}

type Iphone struct{

}

func (iphone Iphone) call(){
	fmt.Println("I am Iphone,I can call you!")
}

func main(){
	
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
//	phone.add()

	var phone1* NokiaPhone = new(NokiaPhone)
	phone1.add()
	
	var phone2* NokiaPhone
	phone2 = new(NokiaPhone)
	phone2.add()

}
