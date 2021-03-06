package testvar

import "fmt"

func SwitchParam(x int, y int) {
	x, y = y, x
	fmt.Printf("x:%d, y:%d\n", x, y)
}

func MultiValue() (str1, str2, str3 string) {
	str1 = "hello"
	str2 = "world"
	return
}

// const可以对标其他语言中的枚举
func TestConst() {
	const Pi float64 = 3.1459265358979323846
	const (
		KB  int64 = 1024
		eof       = -1
	)
	const name1 = iota
	const name2 = iota
	const name3 = iota
	fmt.Printf("Pi:%f, KB:%d, name1:%d, name2:%d, name3:%d\n", Pi, KB, name1, name2, name3)
	const (
		ASC  = iota
		DESC = iota
	)
	number := int(10)
	fmt.Println("number:", number)
	var data = [2]uint8{0x1, 0x2}
	fmt.Printf("data:%d\n", data)
	var str = "hello"[:2]
	fmt.Printf("first letter is %s\n", str)
}

//字符串遍历过程中每一个字符类型是rune
func TestString() {
	var data string = "hello世界"
	for index, wd := range data {
		fmt.Printf("index %d, word: %c\n", index, wd)
	}
}

func TestConvert() {
	//var flag bool
	//flag = bool(1) 布尔值是不兼容类型
	var char int8 = 'a'
	var data int = int(char)
	fmt.Println("data:", data)
}

func TestAlias() {
	type mytype int32
	var data mytype = 100
	fmt.Printf("my type is:%T\n", data)
}
