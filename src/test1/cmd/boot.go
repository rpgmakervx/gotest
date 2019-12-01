package main

import (
	"fmt"
	"os"
	"strconv"
)
import "test1/cal"

func main() {

	args := os.Args[1:]
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	switch args[0] {
	case "+":
		if len(args) < 3 {
			fmt.Println("Usage calc + <int> <int>")
			return
		}
		x, err1 := strconv.ParseFloat(args[1], 64)
		y, err2 := strconv.ParseFloat(args[2], 64)
		if err1 != nil || err2 != nil {
			handleError("+")
		}
		result := cal.Add(x, y)
		fmt.Println(result)
	case "-":
		if len(args) < 3 {
			handleError("-")
			return
		}
		x, err1 := strconv.ParseFloat(args[1], 64)
		y, err2 := strconv.ParseFloat(args[2], 64)
		if err1 != nil || err2 != nil {
			handleError("-")
			return
		}
		result := cal.Sub(x, y)
		fmt.Println(result)
	case "*":
		if len(args) < 3 {
			handleError("*")
			return
		}
		x, err1 := strconv.ParseFloat(args[1], 64)
		y, err2 := strconv.ParseFloat(args[2], 64)
		if err1 != nil || err2 != nil {
			handleError("*")
			return
		}
		result := cal.Multiply(x, y)
		fmt.Println(result)
	case "/":
		if len(args) < 3 {
			handleError("/")
			return
		}
		x, err1 := strconv.ParseFloat(args[1], 64)
		y, err2 := strconv.ParseFloat(args[2], 64)
		if err1 != nil || err2 != nil {
			handleError("/")
			return
		}
		result := cal.Div(x, y)
		fmt.Println(result)
	default:
		Usage()
	}
}

func Usage() {
	fmt.Println("Usage: cal command [args]...")
	fmt.Println("command here:")
	fmt.Println("+: 加法")
	fmt.Println("-: 减法")
	fmt.Println("*: 乘法")
	fmt.Println("/: 除法")
	fmt.Println("/: 除法")
}

func handleError(op string) {
	fmt.Println("Usage calc " + op + " <int> <int>")
}
