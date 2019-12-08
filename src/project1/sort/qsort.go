package sort

import "fmt"

func QSort(array []int) {
	partition(array, 0, len(array)-1)
}

func partition(array []int, low int, high int) {
	if low >= high {
		return
	}
	l := low
	h := high
	for l < h {
		for array[l] > array[h] && l < h {
			fmt.Printf("highpart-skip num: array[%d]=%d, array[%d]=%d\n", l, array[l], h, array[h])
			h--
		}
		array[l], array[h] = array[h], array[l]
		fmt.Println("high-part low:", low, " high:", high, " l:", l, " h:", h, " lv:", array[l], "hv:", array[h], "array:", array)
		for array[l] > array[h] && l < h {
			fmt.Printf("lowpart-skip num: array[%d]=%d, array[%d]=%d\n", l, array[l], h, array[h])
			l++
		}
		array[l], array[h] = array[h], array[l]
		fmt.Println("low-part low:", low, " high:", high, " l:", l, " h:", h, " lv:", array[l], "hv:", array[h], "array:", array)
	}
	fmt.Println("term-part low:", low, " high:", high, " l:", l, " h:", h, " array:", array)
	fmt.Println("----------finish this term-----------")
	partition(array, low, l-1)
	partition(array, h+1, high)
}
