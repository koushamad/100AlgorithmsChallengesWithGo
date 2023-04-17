package NightRoute

const PerMile = 15

func solution(city [][]int) int {
	return recursive(city, 0)
}

func recursive(city [][]int, step int) int {
	lenCity := len(city)

	if step == lenCity-1 {
		return 0
	}

	result := make([]int, lenCity)
	for i := 0; i < lenCity; i++ {
		result[i] = -1
	}

	for i := step + 1; i < lenCity; i++ {
		cost := city[step][i]
		if cost != -1 {
			result[i] = cost
			result[i] += recursive(city, i)
		}
	}

	return min(result)
}

func min(items []int) int {
	min := 100

	for _, v := range items {
		if v == -1 {
			continue
		}
		if v <= min {
			min = v
		}
	}

	return min
}
