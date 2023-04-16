package DataStructure

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
