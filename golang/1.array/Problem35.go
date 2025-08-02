package main

import "sort"

func SmallerThanCurrentNumber(nums []int) (ret []int) {
	nc := append([]int(nil), nums...)
	sort.Ints(nc)

	m := make(map[int]int)

	for i, v := range nc {
		_, e := m[v]
		if !e {
			m[v] = i
			continue
		}
	}

	ret = make([]int, len(nums))
	for i, v := range nums {
		sc, e := m[v]
		if e {
			ret[i] = sc
			continue
		}
		ret[i] = -1
	}

	return ret
}
