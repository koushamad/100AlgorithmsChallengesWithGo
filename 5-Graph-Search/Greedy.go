package GraphSearch

import "sort"

func FindMinArrowShots(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}
	// sort the points by left, order by right
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})
	// store the current target range (left, right) for arrow
	cnt := 0
	left := points[0][0]
	right := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > right {
			// if the current ballon is not in the range of (left, right)
			// fire an arrow to the target range to hit the group of ballons within the range (left, right)
			cnt++
			// update the range for the next arrow
			left = points[i][0]
			right = points[i][1]
		} else {
			// if current ballon is within the range of (left, right)
			// update the range if needed
			left = findMax(left, points[i][0])
			right = findMin(right, points[i][1])
		}
	}
	// +1: at the end, fire the last arrow to hit the last group of ballons
	return cnt + 1
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
