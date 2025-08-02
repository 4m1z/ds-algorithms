package main

// o(n)
func containDuplicate(nums []int) bool {
	f := false

	m := make(map[int]bool)

	for _, v := range nums {
		_, exsits := m[v]
		if exsits {
			f = true
			break
		}
		m[v] = true
	}

	return f
}
