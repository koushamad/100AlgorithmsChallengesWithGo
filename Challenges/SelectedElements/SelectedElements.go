package SelectedElements

import (
	"fmt"
	ElementarySorts "github.com/koushamad/100AlgorithmsChallengesWithGo/4-Elementary-Sorts"
)

// solution QuickSort dynamic programming O(N*(P+2Q+3R))
func solution(arr []int, p, q, r int) int {
	sortedArr := ElementarySorts.QuickSort(arr)
	fmt.Println(sortedArr)
	sum := 0

	for _, num := range sortedArr[:p+q*2+r*3] {
		sum += num
	}

	return sum
}
