package testtime

import (
	"fmt"
	"time"
)

func TestTime() {
	//testNow()
	//testDuration()
	testFormat()
}

func testNow() {
	var now = time.Now()
	fmt.Println("now time is:", now)
	fmt.Printf("%02d.%02d.%4d\n", now.Day(), now.Month(), now.Year())

}

func testDuration() {
	var now = time.Now()
	week := 60 * 60 * 24 * 7 * 1e9
	newNow := now.Add(time.Duration(week))
	fmt.Println("newNow time is:", newNow)
}

func testFormat() {
	var now = time.Now()
	//2006-01-02 15:04:05这个时间比较魔性，貌似是golang诞生的时间。
	//所以不同于java/python等使用 yyyMMdd的方式，需要用这个时间的格式来处理格式化时间
	var format = now.Format("2006-01-02 15:04:05")
	fmt.Println("origin now:", now, " format now:", format)
}
