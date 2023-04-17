package FancyRide

const Off = 20.0

func solution(l int, fares []float64) string {
	cars := []string{"UberX", "UberXI", "UberPlus", "UberBlack", "UberSUV"}

	for i, c := range fares {
		cost := c * float64(l)
		if cost > Off {
			return cars[i-1]
		}
	}

	return cars[4]
}
