package main

import (
	"fmt"
)

type I interface {
	A()
}

type impl struct{}

func (i *impl) A() {}

func isNil[T any, PT *T](i PT) bool { //HL12
	var empty PT
	return i == empty
}

func main() {

	a := &impl{}
	var b I
	c := (*impl)(nil)
	// START OMIT
	fmt.Printf("isNil(%v) %t \n", a, isNil(a))
	fmt.Printf("isNil(%v) %t \n", b, isNil(b))
	fmt.Printf("isNil(%v) %t \n", c, isNil(c))
	fmt.Printf("%v == nil %t \n", a, a == nil)
	fmt.Printf("%v == nil %t \n", b, b == nil)
	fmt.Printf("%v == nil %t \n", c, c == nil)
	//END OMIT

}
