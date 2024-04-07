package main

func pivotIndex1(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	var sum int
	var sumArr = []int{}

	for _, n := range nums {
		sum += n
		sumArr = append(sumArr, sum)
	}

	for i, s := range nums {

		var sumI int

		if i == 0 {
			sumI = 0
		} else {
			sumI = sumArr[i-1]
		}

		if sum-sumI*2 == s {
			return i
		}
	}

	return -1
}
