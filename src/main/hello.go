package main

import (
	"fmt"
)

func main() {
	//fmt.Println("this is code4j's first go programe!")
	//a, b := addAll(1, 2)
	//fmt.Printf("a and b add 1 is:%d, %d\n", a, b)
	//testFunction()
	//flyer1 := new(Bird)
	//flyer2 := new(Eagle)
	//flyer1.fly()
	//flyer2.fly()
	var numbers []int = [] int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var channel = make(chan int, 2);
	sum(numbers[:len(numbers) / 2], channel)
	sum(numbers[len(numbers) / 2:], channel)
	fmt.Println("sum is ", <-channel, <-channel)
}

func addAll(a int, b int) (int, int) {
	return a + 1, 1 + b
}

func testFunction() {
	f := func(x int, y int) int {
		return x + y
	}

	fmt.Println(f(1, 2))
}

type Bird struct {
}

type Eagle struct {
}

type IFly interface {
}

func (bird *Bird) fly() {
	fmt.Println("bird fly")
}

func (bird *Eagle) fly() {
	fmt.Println("eagle fly")
}

func sum(numbers [] int, channel chan int) {
	var sum int = 0;
	for _, val := range numbers {
		sum += val
	}
	channel <- sum
}
