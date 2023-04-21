package FirstAndLastPosition

func solution(arr []int, target int) []int {
	result := make([]int, 2)
	if len(arr) == 0 || arr[0] > target || arr[len(arr)-1] < target {
		return []int{-1, -1}
	}

	result[0] = findFirst(arr, target)
	result[1] = findLast(arr, target)

	return result
}

func findFirst(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		midVal := arr[mid]

		if midVal == target && arr[mid-1] < target {
			return mid
		} else if midVal < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func findLast(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		midVal := arr[mid]

		if midVal == target && arr[mid+1] > target {
			return mid
		} else if midVal <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
