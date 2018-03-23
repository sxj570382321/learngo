package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Conn
}

type Conn interface {
	connector()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) connector() {
	fmt.Printf("connected:%s!\n", pc.name)
}

type empty interface {
}

func main() {
	var in USB
	in = PhoneConnector{"PhoneConnector"}
	//	in = new(PhoneConnector{"PhoneConnector"})
	in.connector()

	Disconnect(in)

	var a empty
	fmt.Println(a == nil)

	var p *int = nil
	a = p
	fmt.Println(a == nil)

	Disconnect1(in)

	var cc Conn
	cc = Conn(in) //USB is child,and Conn is father,child is father,but father is not child
	cc.connector()

	/*
		var in1 USB
		var cc1 Conn
		in1 = USB(cc1) //but father is not child
		in1.connector()
	*/
}

//ok-button
func Disconnect(sa USB) {
	if pc, ok := sa.(PhoneConnector); ok {
		fmt.Printf("Disconnected:%s!\n", pc.name)
		return
	}

	fmt.Println("Unknown device!")
}

//type-switch
func Disconnect1(sa empty) {
	switch i := sa.(type) {
	case PhoneConnector:
		fmt.Printf("Disconnected1:%s!\n", i.name)
	default:
		fmt.Println("Unknown device!")
	}

}
