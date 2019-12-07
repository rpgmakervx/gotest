package testfunc

import (
	"errors"
	"fmt"
)

func TestFunction(x int32, y int32) (int32, error) {
	if y == 0 {
		err := errors.New("illegal argument exception, y cannot be zero")
		return 0, err
	}
	return x / y, nil
}
func TestSameTypeParam(x, y, z int32) (int32, int32, error) {
	return x + y + z, 0, nil
}

func TestAnonymous() {
	f := func() {
		fmt.Println("hello anonymous")
	}
	f()
	f2 := func(x int32, y int32) int32 {
		return x + y
	}
	ret := f2(1, 2)
	fmt.Printf("%d + %d = %d\n", 1, 2, ret)

	code, msg := func(x, y int32) (code int32, msg string) {
		msg = "x + y = " + string(x+y)
		code = 200
		return
	}(10, 20)
	fmt.Printf("code:%d, message:%s\n", code, msg)

	var a = 10
	var b = 100

	func() {
		a, b = b, a
		fmt.Printf("anonymous a:%d, b:%d\n", a, b)
	}()
	fmt.Printf("out a:%d, b:%d\n", a, b)
}

func TestClosure() {
	f := Closure()
	fmt.Println("------test closure-----")
	fmt.Printf("x * x = %d\n", f())
	fmt.Printf("x * x = %d\n", f())
	fmt.Printf("x * x = %d\n", f())
	fmt.Printf("x * x = %d\n", f())
}

func Closure() func() int32 {
	var x int32
	return func() int32 {
		x++
		return x * x
	}
}
