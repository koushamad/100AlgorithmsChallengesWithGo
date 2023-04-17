package FareEstimator

func solution(rt int, rd int, cpm []float64, cpk []float64) []float64 {
	lecpm := len(cpm)
	result := make([]float64, lecpm)

	for i := 0; i < lecpm; i++ {
		result[i] = cpm[i]*float64(rt) + cpk[i]*float64(rd)
	}

	return result
}
