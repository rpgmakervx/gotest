package teststring

import (
	"fmt"
	"strconv"
	"strings"
)

func TestString() {
	testPrSuf()
	testContains()
	testIndex()
	testReplaceAndCountAndRepeat()
	testLowUp()
	testSplitAndJoin()
	testTransfer()
}

func testPrSuf() {
	var str = "index.json"
	if strings.HasPrefix(str, "index") {
		fmt.Println("this is an index file")
	}
	if strings.HasSuffix(str, "json") {
		fmt.Println("this is a config file")
	}
}

func testContains() {
	var str = "My name is xingtianyu"
	if strings.Contains(str, "name") {
		fmt.Println("this is an introduction")
	}
}

func testIndex() {
	var str = "x + y = z"
	index := strings.Index(str, "=")
	fmt.Println("equals index:", index)

	str = "xyz abc x+y"
	index = strings.LastIndex(str, "x")
	fmt.Println("last index:", index)
}

func testReplaceAndCountAndRepeat() {
	var str = "index.json"
	result := strings.ReplaceAll(str, "json", "log")
	fmt.Println("origin:", str, " replace:", result)

	var str2 = "abc acb bac bca"
	var count = strings.Count(str2, "ac")
	fmt.Println("count:", count)

	var str3 = strings.Repeat(str, 10)
	fmt.Println("repeat:", str3)
}

func testSplitAndJoin() {
	var str1 = "xxx|yyy|zzz"
	var str2 = "-"
	segments := strings.Split(str1, "|")
	fmt.Println("string segment:", segments)
	var result = strings.Join(segments, str2)
	for index, s := range result {
		fmt.Println("index:", index, " word:", s)
	}
	fmt.Println("contact:", result)

	var str3 = "I have a dream"
	segments = strings.Fields(str3)
	fmt.Println("segment fields:", segments)
}

func testLowUp() {
	var str = " I have a dream that on day I will go to the UK"
	var upper = strings.ToUpper(str)
	var lower = strings.ToLower(upper)
	fmt.Println("origin", str, "upper:", upper)
	fmt.Println("origin", str, "lower:", lower)
}

func testTransfer() {
	var str = "10"
	num, _ := strconv.Atoi(str)
	fmt.Printf("num convert from str:%d, %T\n", num, num)
	str = strconv.Itoa(num)
	fmt.Printf("str convert from num:%s, %T\n", str, str)
	var fstr = "10.5"
	fnum, _ := strconv.ParseFloat(fstr, 32)
	fmt.Printf("fstr convert from fnum:%s, %T\n", fstr, fnum)
}
