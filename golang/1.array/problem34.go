package main

func FindDisappearedNumbers(nums []int) []int {
	mn := []int{}
	s := make(map[int]int)
	n := len(nums)

	for i, v := range nums {
		s[v] = i
	}

	for i := 1; i <= n; i++ {
		_, exsits := s[i]
		if !exsits {
			mn = append(mn, i)
		}
	}

	return mn
}
