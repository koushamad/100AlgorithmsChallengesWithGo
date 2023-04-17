package ClosestLocation

import (
	"math"
	"strings"
)

func solution(address string, objects [][]int, names []string) string {
	indexes := make([]int, 0)

	for i, name := range names {
		if CheckAddress(address, name) {
			indexes = append(indexes, i)
		}
	}

	minIndex := -1
	minDist := 200

	for index, object := range objects {
		if len(object) == 2 {
			dist := distPoint(object)
			if dist < minDist {
				minDist = dist
				minIndex = index
			}
		} else {
			dist := distSegment(object)
			minDist = dist
			minIndex = index
		}
	}

	return names[minIndex]
}

func distPoint(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
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

	return int(math.Min(float64(distPoint(segment[:2])), float64(distPoint(segment[2:]))))
}

func distSegment(segment []int) int {
	// Segment is vertical
	if segment[0] == segment[2] {
		if segment[1] >= 0 && segment[3] <= 0 {
			return segment[0] * segment[0]
		}
	} else {
		if segment[1] >= 0 && segment[3] <= 0 {
			return segment[1] * segment[1]
		}
	}

	return int(math.Min(float64(distPoint(segment[:2])), float64(distPoint(segment[2:]))))
}

func CheckAddress(word, name string) bool {
	word = strings.ToLower(word)
	name = strings.ToLower(word)

	// check prefix
	if word == name[:len(word)] {
		return true
	}

	if CheckDiff(word, name) {
		return true
	}

	if CheckMiss(word, name[:len(word)+1]) {
		return true
	}

	if CheckMiss(name[:len(word)-1], word) {
		return true
	}

	return false
}

func CheckDiff(word, name string) bool {
	diff := false
	lenWord := len(word)
	lenName := len(name)

	if lenWord > lenName {
		return false
	}

	for i := 0; i < lenWord; i++ {
		if word[i] != name[i] {
			if diff {
				return false
			}
			diff = true
		}
	}

	return true
}

func CheckMiss(word, name string) bool {
	lenWord := len(word)
	lenName := len(name)
	miss := false
	i, j := 0, 0

	if lenWord+1 != lenName {
		return false
	}

	for i < lenWord && j < lenName {
		if word[i] != name[i] {
			if miss {
				return false
			}

			miss = true
			j++
			continue
		}
		i++
		j++
	}

	return true
}
