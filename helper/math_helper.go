package helper

import "sort"

func Median(nums []int) float64 {
	n := len(nums)
	if n == 0 {
		return 0 // or you can handle it by returning an error or a sentinel value
	}

	// Sort the slice
	sort.Ints(nums)

	if n%2 == 1 {
		// If odd, return the middle element
		return float64(nums[n/2])
	} else {
		// If even, return the average of the two middle elements
		return float64(nums[n/2-1]+nums[n/2]) / 2.0
	}
}
