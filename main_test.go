package main

import (
	"reflect"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	cases := map[string]struct {
		input        []int
		wantSubSlice []int
		wantSum      int
	}{
		"Slice 1": {input: []int{1, 3, 9, 5, 7, -10, 8, -3, 30, -1}, wantSubSlice: []int{1, 3, 9, 5, 7, -10, 8, -3, 30}, wantSum: 50},
		"Slice 2": {input: []int{-21, 3, 76, -95, 5, -33}, wantSubSlice: []int{3, 76}, wantSum: 79},
		"Slice 3": {input: []int{-666, -55}, wantSubSlice: []int{}, wantSum: 0},
		"Slice 4": {input: []int{1, 2, 3, 4}, wantSubSlice: []int{1, 2, 3, 4}, wantSum: 10},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			slice, sum := MaxSubArray(tc.input)
			if !reflect.DeepEqual(slice, tc.wantSubSlice) {
				t.Errorf("MaxSubArray() slice = %v, want %v", slice, tc.wantSubSlice)
			}
			if sum != tc.wantSum {
				t.Errorf("MaxSubArray() sum = %v, want %v", sum, tc.wantSum)
			}
		})
	}
}
