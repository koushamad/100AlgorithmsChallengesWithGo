package PerfectCity

import "math"

func solution(dep, des []float64) float64 {
	rx := math.Abs(des[0] - dep[0])
	ry := math.Abs(des[1] - dep[1])

	if int64(des[0]) == int64(dep[0]) {
		rx1 := math.Ceil(des[0]) - des[0] + math.Ceil(dep[0]) - dep[0]
		rx2 := des[0] - math.Floor(des[0]) + dep[0] - math.Floor(dep[0])
		rx = math.Min(rx1, rx2)
	}

	if int64(des[1]) == int64(dep[1]) {
		ry1 := math.Ceil(des[1]) - des[1] + math.Ceil(dep[1]) - dep[1]
		ry2 := des[1] - math.Floor(des[1]) + dep[1] - math.Floor(dep[1])
		ry = math.Min(ry1, ry2)
	}

	result := rx + ry
	return result
}
