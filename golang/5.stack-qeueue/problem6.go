package main

func Problem6(arr []int) []int {
	res := make([]int, len(arr))

	stk := []int{}

	for i := len(arr) - 1; i >= 0; i-- {
		for len(stk) > 0 && stk[len(stk)-1] <= arr[i] {
			stk = stk[:len(stk)-1]
		}

		if len(stk) == 0 {
			res[i] = -1
		} else {
			res[i] = stk[len(stk)-1]
		}

		stk = append(stk, arr[i])
	}

	return res
}
