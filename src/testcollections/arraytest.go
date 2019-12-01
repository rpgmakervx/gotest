package main

import "fmt"

func main() {
	testSlice()
}

func testArray() {
	week := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	time := [12]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for index, w := range week {
		fmt.Printf("the %d day is %s, addr:%p\n", index, w, &week[index])
	}
	fmt.Println("--------")
	for index, t := range time {
		fmt.Printf("time is %d, addr:%p\n", t, &time[index])
	}
}
func testSlice() {
	array1 := [11]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	array2 := make([]int32, 12)
	array2 = array1[:]
	fmt.Println("array2 is ", array2, "cap:", cap(array2), "len is;", len(array2))
	array3 := append(array2, 12, 13, 14, 15)
	fmt.Println("array3 is ", array3, "cap:", cap(array3), "len is;", len(array3))
	array4 := []int32{16, 17, 18, 19, 20}
	array5 := []int32{21, 22, 23, 24}
	array4 = append(array4, array5...)
	fmt.Println("array4 is ", array4, "cap:", cap(array4), "len is;", len(array4))
	array6 := append(array3, array4...)
	fmt.Println("array6 is ", array6, "cap:", cap(array6), "len is;", len(array6))

	fmt.Println("--------")

	slice1 := []int32{1, 2, 3, 4, 5, 6}
	slice2 := []int32{11, 12, 13}
	copy(slice1, slice2)
	fmt.Println("s1", slice1, "s2", slice2)
	copy(slice2, slice1[3:])
	fmt.Println("s1", slice1, "s2", slice2)

}
