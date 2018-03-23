package main

import (
	"fmt"
)

type TZ int

func (tz *TZ) Increase(m int) {
	*tz = TZ(m)
}

func main() {
	var a TZ
	a.Increase(100)
	fmt.Println(a)
}
