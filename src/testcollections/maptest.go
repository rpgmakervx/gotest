package testcollections

import "fmt"

func TestMap() {
	var dataMap map[int32]string
	dataMap = make(map[int32]string)
	dataMap[1] = "xingtianyu"
	dataMap[2] = "code4j110"
	dataMap[3] = "lalala"
	fmt.Println("dataMap:", dataMap)
	delete(dataMap, 3)
	fmt.Println("dataMap:", dataMap)
	name, got := dataMap[1]
	fmt.Println(name, got)
	var null string
	dataMap[4] = null

	if a := 0; a <= 0 {
		fmt.Println(a)
		return
	}
	fmt.Println("a is over 0")
}
