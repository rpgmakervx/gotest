package pointer

import "fmt"

func TetsPointer() {
	var x int32 = 10
	var pointer = &x
	fmt.Printf("p = %v, &x = %v\n", pointer, &x)
	*pointer = 11
	fmt.Printf("p = %v, x = %d\n", pointer, x)

	var num int32 = 208
	var p *int32
	p = new(int32)
	p = &num
	fmt.Printf("p = %v, value = %d\n", p, *p)

	var a int32 = 10
	var b int32 = 100
	pointer = &a
	p = &b
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("pointer = %p, p = %p\n", pointer, p)
	Swap(pointer, p)
	fmt.Printf("a = %d, b = %d\n", a, b)
	SwapAddr(pointer, p)
	fmt.Printf("pointer = %p, p = %p\n", pointer, p)
}

func Swap(x *int32, y *int32) {
	*x, *y = *y, *x
}
func SwapAddr(x *int32, y *int32) {
	x, y = y, x
	fmt.Printf("pointer = %p, p = %p\n", x, y)
}
