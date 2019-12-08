package sort

func BubbleSort(array []int) {
	for index := 1; index < len(array); index++ {
		for idx := index - 1; idx < len(array); idx++ {
			if array[idx] > array[idx-1] {
				array[idx], array[idx-1] = array[idx-1], array[idx]
			}
		}
	}
}
