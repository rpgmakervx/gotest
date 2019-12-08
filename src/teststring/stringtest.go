package teststring

import (
	"fmt"
	"strings"
)

func TestString() {
	testPrSuf()
	testContains()
	testIndex()
	testReplaceAndCountAndRepeat()
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
}
