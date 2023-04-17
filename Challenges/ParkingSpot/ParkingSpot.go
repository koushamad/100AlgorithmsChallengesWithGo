package ParkingSpot

func solution(carDimensions []int, parkingLot [][]int, luckySpot []int) bool {
	result := false

	if !IsEmpty(parkingLot, luckySpot[:2], luckySpot[2:]) {
		return false
	}

	if isVertical(carDimensions) {
		l := left(parkingLot, luckySpot[:2], luckySpot[2:])
		r := right(parkingLot, luckySpot[:2], luckySpot[2:])
		result = l || r
	} else {
		u := up(parkingLot, luckySpot[:2], luckySpot[2:])
		d := down(parkingLot, luckySpot[:2], luckySpot[2:])
		result = u || d
	}

	return result
}

func IsEmpty(parkingLot [][]int, p1, p2 []int) bool {

	for i := p1[0]; i <= p2[0]; i++ {
		for j := p1[1]; j <= p2[1]; j++ {
			if parkingLot[i][j] == 1 {
				return false
			}
		}
	}

	return true
}

func InSpace(parkingLot [][]int, p1, p2 []int) bool {
	lenX := len(parkingLot) - 1
	lenY := len(parkingLot[0]) - 1

	if p1[0] > lenX || p1[0] < 0 || p2[0] > lenX || p2[0] < 0 {
		return false
	}
	if p1[1] > lenY || p1[1] < 0 || p2[1] > lenY || p2[1] < 0 {
		return false
	}

	return true
}

func right(parkingLot [][]int, p1, p2 []int) bool {
	tp1 := make([]int, 2)
	tp2 := make([]int, 2)

	result := false

	for InSpace(parkingLot, p1, p2) {
		tp1[0] = p1[0]
		tp1[1] = p1[1] + 1
		tp2[0] = p2[0]
		tp2[1] = p2[1] + 1

		if IsEmpty(parkingLot, p1, p2) == false {
			result = false
			break
		} else {
			result = true
			p1[0] = tp1[0]
			p1[1] = tp1[1]
			p2[0] = tp2[0]
			p2[1] = tp2[1]
		}
	}

	return result
}

func left(parkingLot [][]int, p1, p2 []int) bool {
	tp1 := make([]int, 2)
	tp2 := make([]int, 2)

	result := false

	for InSpace(parkingLot, p1, p2) {
		tp1[0] = p1[0]
		tp1[1] = p1[1] - 1
		tp2[0] = p2[0]
		tp2[1] = p2[1] - 1

		if IsEmpty(parkingLot, p1, p2) == false {
			result = false
			break
		} else {
			result = true
			p1[0] = tp1[0]
			p1[1] = tp1[1]
			p2[0] = tp2[0]
			p2[1] = tp2[1]
		}
	}

	return result
}

func down(parkingLot [][]int, p1, p2 []int) bool {
	tp1 := make([]int, 2)
	tp2 := make([]int, 2)

	result := false

	for InSpace(parkingLot, p1, p2) {
		tp1[0] = p1[0] + 1
		tp1[1] = p1[1]
		tp2[0] = p2[0] + 1
		tp2[1] = p2[1]

		if IsEmpty(parkingLot, p1, p2) == false {
			result = false
			break
		} else {
			result = true
			p1[0] = tp1[0]
			p1[1] = tp1[1]
			p2[0] = tp2[0]
			p2[1] = tp2[1]
		}
	}

	return result
}

func up(parkingLot [][]int, p1, p2 []int) bool {
	tp1 := make([]int, 2)
	tp2 := make([]int, 2)

	result := false

	for InSpace(parkingLot, p1, p2) {
		tp1[0] = p1[0] - 1
		tp1[1] = p1[1]
		tp2[0] = p2[0] - 1
		tp2[1] = p2[1]

		if IsEmpty(parkingLot, p1, p2) == false {
			result = false
			break
		} else {
			result = true
			p1[0] = tp1[0]
			p1[1] = tp1[1]
			p2[0] = tp2[0]
			p2[1] = tp2[1]
		}
	}

	return result
}

func isVertical(p []int) bool {
	return p[0] > p[1]
}
