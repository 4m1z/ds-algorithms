package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestProblem28(t *testing.T) {
	result := Problem28([]int{-1, 2, -2, 1, -1, 2})
	expectedResult := [][]int{
		{-1, -1, 2},
	}

	// Helper function to sort a 2D slice
	sort2D := func(data [][]int) {
		for i := range data {
			sort.Ints(data[i]) // Sort inner slices
		}
		sort.Slice(data, func(i, j int) bool {
			for k := range data[i] {
				if k >= len(data[j]) {
					return false
				}
				if data[i][k] != data[j][k] {
					return data[i][k] < data[j][k]
				}
			}
			return len(data[i]) < len(data[j])
		})
	}

	sort2D(result)
	sort2D(expectedResult)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("**first test case ** Expected (%v) is not same as"+
			"actual (%v)", result, expectedResult)
	}
}
