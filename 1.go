package main

import "fmt"

func quick(left, right int, arr []int) {
	l, r := left, right
	mid := arr[(left+right)/2]
	for r > l {
		for mid < arr[r] {
			r--
		}
		for mid > arr[l] {
			l++
		}
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		if mid == arr[r] {
			l++
		}
		if mid == arr[l] {
			r--
		}
	}
	if l == r {
		r--
		l++
	}
	if r > left {
		quick(left, r, arr)
	}
	if l < right {
		quick(l, right, arr)
	}
}
func main() {
	arr := []int{4, 2, 5, 7, 6, 3, 2, 9, 1}
	//for i := 0; i < len(arr); i++ {
	//	for j := i + 1; j < len(arr); j++ {
	//		if arr[j] < arr[i] {
	//			arr[j], arr[i] = arr[i], arr[j]
	//		}
	//	}
	//}
	quick(0, len(arr)-1, arr)
	fmt.Println(arr)
}
