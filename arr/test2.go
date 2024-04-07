package main

func pivotIndex2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	var sum int

	for _, n := range nums {
		sum += n
	}

	var sumLeft int
	if sumLeft == sum-nums[0] {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		sumLeft += nums[i-1]
		if sumLeft*2 == sum-nums[i] {
			return i
		}
	}

	return -1
}
