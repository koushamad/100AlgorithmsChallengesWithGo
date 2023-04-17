package ConvexHull

import (
	"fmt"
	"sort"
)

type Point struct {
	X, Y int
}
type Points []Point

func (pt Points) Len() int      { return len(pt) }
func (pt Points) Swap(a, b int) { pt[a], pt[b] = pt[b], pt[a] }
func (pt Points) Less(a, b int) bool {
	if pt[a].X < pt[b].X {
		return true
	} else if pt[a].X > pt[b].X {
		return false
	}
	return pt[a].Y < pt[b].Y
}

func solution(points []Point) []Point {

	if points == nil || len(points) < 3 {
		return points
	}
	sort.Sort(Points(points))

	var lower []Point
	var upper []Point
	fmt.Println("sorted", points)
	for _, p := range points {
		n := len(lower)
		for n > 1 && Ccw(lower[n-2], lower[n-1], p) < 0 {
			lower = lower[:n-1]
			n = len(lower)
		}
		lower = append(lower, p)
	}

	sort.Sort(sort.Reverse(Points(points)))
	fmt.Println("revered", points)
	for _, p := range points {
		n := len(upper)
		for n > 1 && Ccw(upper[n-2], upper[n-1], p) < 0 {
			upper = upper[:n-1]
			n = len(upper)
		}
		upper = append(upper, p)
	}
	_map := make(map[Point]bool)
	rst := append(lower[:len(lower)-1], upper[:len(upper)-1]...)

	for i, p := range rst {
		if _, ok := _map[p]; ok {
			rst = append(rst[:i], rst[i+1:]...)
		} else {
			_map[p] = true
		}

	}
	return rst
}

func Ccw(a, b, c Point) int {
	return (b.X-a.X)*(c.Y-a.Y) - (c.X-a.X)*(b.Y-a.Y)
}
