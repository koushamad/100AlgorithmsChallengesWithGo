package DataStructure

import "math"

func getKeyArray(y, x int) int {
	return y*10 + x*1
}

func getKeyIndexesArray(matrix [][]int) map[int]int {
	indexes := make(map[int]int)
	lenY := len(matrix)
	lenX := len(matrix[0])

	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			if matrix[y][x] == 0 {
				continue
			}

			key := getKeyArray(y, x)
			indexes[key] = matrix[y][x]
		}
	}

	return indexes
}

func DistPoint(p1, p2 []int) int {
	return int(math.Abs(float64(p1[0]-p2[0]))) * int(math.Abs(float64(p1[1]-p2[1])))
}

func DistSegment(point []int, segment []int) int {
	// Segment is vertical
	if segment[0] == segment[2] && point[1] >= int(math.Min(float64(segment[1]), float64(segment[3]))) && point[1] <= int(math.Max(float64(segment[1]), float64(segment[3]))) {
		return int(math.Abs(float64(point[0] - segment[0])))
	} else if segment[1] == segment[3] && point[0] >= int(math.Min(float64(segment[0]), float64(segment[2]))) && point[0] <= int(math.Max(float64(segment[0]), float64(segment[2]))) {
		return int(math.Abs(float64(point[1] - segment[1])))
	}

	return int(math.Min(float64(DistPoint(point, segment[:2])), float64(DistPoint(point, segment[2:]))))
}
