package main

import "fmt"

func MaxSubArray(arr []int) ([]int, int) {
	var max, startIndex, endIndex, sum, startSum int

	for index, val := range arr {
		sum = sum + val
		switch {
		case sum > max:
			max = sum
			startIndex = startSum
			endIndex = index + 1
			fmt.Println("index:", index, "value:", val, "sum:", sum, "> max", max, "| subslice", arr[startIndex:endIndex], "sum:", sum)
		case sum <= 0:
			sum = 0
			startSum = index + 1
			fmt.Println("index:", index, "value:", val, "| sum:", sum, "< 0", "| subslice", arr[startIndex:endIndex], "sum:", sum)
		}
	}
	return arr[startIndex:endIndex], max
}
