package main

import "reflect"

func partitionFn(arr []int, left, right int) int {
	pivot := arr[right]
	partitionIndex := left - 1
	swapFn := reflect.Swapper(arr)

	for i := left; i < right; i++ {
		if pivot > arr[i] {
			partitionIndex++
			swapFn(i, partitionIndex)
		}
	}

	swapFn(partitionIndex+1, right)
	return partitionIndex + 1
}

func quickSort(arr []int, left, right int) { // o(2^N)
	if left >= right {
		return
	}

	partionIndex := partitionFn(arr, left, right)

	quickSort(arr, left, partionIndex-1)
	quickSort(arr, partionIndex+1, right)
}

func quickSort2(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[len(arr)/2]
	var less, equal, greater []int
	for _, value := range arr {
		switch {
		case value < pivot:
			less = append(less, value)
		case value == pivot:
			equal = append(equal, value)
		case value > pivot:
			greater = append(greater, value)
		}
	}
	quickSort2(less)
	quickSort2(greater)
	copy(arr, append(append(less, equal...), greater...))
}
