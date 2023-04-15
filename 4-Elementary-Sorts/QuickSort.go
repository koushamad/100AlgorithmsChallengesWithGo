package ElementarySorts

// QuickSort ~NLogN Not Stable
type quickSort struct {
	List []int
}

func QuickSort(list []int) []int {
	q := &quickSort{
		List: list,
	}
	q.sort(0, len(list)-1)
	return q.List
}

func (q *quickSort) sort(lo, hi int) {
	if lo < hi {
		p := q.partition(lo, hi)
		q.sort(lo, p-1)
		q.sort(p+1, hi)
	}
}

func (q *quickSort) partition(lo, hi int) int {
	pivot := q.List[hi]
	i := lo

	for j := lo; j < hi; j++ {
		if q.List[j] < pivot {
			q.List[i], q.List[j] = q.List[j], q.List[i]
			i++
		}
	}

	q.List[i], q.List[hi] = q.List[hi], q.List[i]

	return i
}
