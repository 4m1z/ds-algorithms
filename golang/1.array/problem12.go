package main

import (
	"sort"
)

// O(n*n)
func CountTriplets(arr []int) int {
	sort.Ints(arr)
	// counter
	cr := 0
	// a + b = c
	for i, _ := range arr {
		if i-1 >= 0 && arr[i-1] == arr[i] {
			continue
		}
		res := isSumPair(arr, arr[i])

		if res {
			cr++
		}
	}

	return cr
}

func isSumPair(arr []int, target int) bool {
	a := 0
	b := len(arr) - 1

	for a < b {
		if arr[a]+arr[b] == target {
			return true
		} else if arr[a]+arr[b] > target {
			b--
		} else {
			// a + b < target
			a++
		}
	}

	return false
}
