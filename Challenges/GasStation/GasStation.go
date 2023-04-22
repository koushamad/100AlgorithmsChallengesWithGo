package GasStation

// O(n^2)
func solution(gas, cost []int) int {
	lenGas := len(gas)

	for i := 0; i < lenGas; i++ {
		if resolve(gas, cost, i) {
			return i
		}
	}

	return -1
}

func resolve(gas, cost []int, index int) bool {
	i := index
	lenGas := len(gas)
	sum := 0
	for sum >= 0 {
		sum += gas[i] - cost[i]

		i++
		if i == lenGas {
			i = 0
		}
		if i == index {
			return true
		}
	}

	return false
}

// O(n)
func solution2(gas, cost []int) int {
	lenGas := len(gas)

	sum := 0
	co := 0
	can := 0
	for i := 0; i < lenGas; i++ {
		sum += gas[i] - cost[i]

		if sum < 0 {
			can = i + 1
			co += sum
			sum = 0
		}
	}

	if sum+co >= 0 {
		return can
	}

	return -1
}
