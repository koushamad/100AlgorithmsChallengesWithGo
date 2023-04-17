package Selection

func solution(nums []int, k int) int {
	return doFindKthLargest(nums, k, 0, len(nums)-1)
}

func doFindKthLargest(nums []int, k int, start, end int) int {
	nlen := len(nums) - 1
	targetPos := nlen - k

	for {
		pivotIndex := partition(nums, start, end)
		if pivotIndex == targetPos {
			return nums[pivotIndex]
		} else if pivotIndex < targetPos {
			start = pivotIndex + 1
		} else {
			end = pivotIndex - 1
		}
	}
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]
	left, right := start+1, end

	for left <= right {
		for left <= right && nums[left] <= pivot {
			left++
		}
		for left <= right && nums[right] > pivot {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[right], nums[start] = nums[start], nums[right]
	return right
}
