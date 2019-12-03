package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	result := rand.Int31n(100)
	for {
		var x int32
		fmt.Println("输入一个0~100的数字:")
		fmt.Scanf("%d", &x)
		if result > x {
			fmt.Println("低了！")
		} else if result < x {
			fmt.Println("高了！")
		} else {
			fmt.Println("回答正确！")
			break
		}
	}
}
