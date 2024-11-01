package parallelbinarysearch

// meteors
func Problem4() {
}

// input := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
// targets := []int{5, 13, 19, 8}
func parallelBinarySearch(arr []int, targets []int) []int {
	res := make([]int, len(targets))

	l := make([]int, len(targets))
	r := make([]int, len(targets))

	for i := range targets {
		l[i] = 0
		r[i] = len(arr) - 1
		res[i] = -1
	}

	for {
		done := true

		for i := range targets {
			if l[i] <= r[i] {
				done = false
				mid := l[i] + (r[i]-l[i])/2
				if arr[mid] == targets[i] {
					res[i] = mid
					l[i] = r[i] + 1
				} else if arr[mid] < targets[i] {
					l[i] = mid + 1
				} else {
					r[i] = mid - 1
				}
			}
		}

		if done {
			break
		}
	}

	return res
}
