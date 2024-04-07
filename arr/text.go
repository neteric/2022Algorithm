package main

func pivotIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	head := 0
	tail := len(nums) - 1

	sumRange := func(start int, end int) int {
		sum := 0

		if start < 0 || end > len(nums)-1 {
			return 0
		}

		for i := start; i <= end; i++ {
			sum += nums[i]
		}

		return sum
	}

	for head != tail {
		if sumRange(0, head-1) >= sumRange(tail+1, len(nums)-1) {
			tail--
		} else {
			head++
		}
	}

	if sumRange(0, head) == sumRange(tail, len(nums)-1) {
		return head
	}

	return -1

}
