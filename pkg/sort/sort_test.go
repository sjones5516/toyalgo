package sort

import (
	"slices"
	"testing"
)

type sortFunc func([]int) []int

var testCases = []struct {
    name     string
    input    []int
    expected []int
}{
    {"already sorted", []int{1, 2, 3}, []int{1, 2, 3}},
    {"unsorted", []int{3, 1, 2}, []int{1, 2, 3}},
    {"reversed", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
    {"with duplicates", []int{4, 2, 2, 3, 1}, []int{1, 2, 2, 3, 4}},
    {"single element", []int{42}, []int{42}},
    {"empty", []int{}, []int{}},
}

func testSortingFunction(t *testing.T, fn sortFunc, fnName string) {
    for _, tc := range testCases {
        t.Run(fnName+"_"+tc.name, func(t *testing.T) {
            result := fn(append([]int{}, tc.input...)) // copy input to avoid mutation
			if !slices.Equal(tc.expected, result) {
				t.Errorf("For %s: expected %v, got %v", tc.name, tc.expected, result)
			}
        })
    }
}

func TestMergeSort(t *testing.T) {
	testSortingFunction(t, MergeSort, "MergeSort")
}