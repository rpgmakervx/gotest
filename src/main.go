package main

import (
	"fmt"
	"testcollections"
	"testifelse"
	"testvar"
)

func main() {
	testvar.SwitchParam(1, 2)
	s1, s2, _ := testvar.MultiValue()
	fmt.Println(s1, s2)
	testvar.TestConst()
	testvar.TestString()
	testcollections.TestArray()
	testcollections.TestMap()
	result := testifelse.TestSwitch(2)
	fmt.Println("result is:", result)
	var x int32
	fmt.Println("please input a variable:")
	fmt.Scanf("%d", &x)
	fmt.Printf("x is %d\n", x)
}
