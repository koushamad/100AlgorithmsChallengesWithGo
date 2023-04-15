package ElementarySorts

// QuickSort ~NLogN Not Stable
type quickSort3Way struct {
	List []int
}

func QuickSort3Way(list []int) []int {
	q := &quickSort{
		List: list,
	}
	q.sort(0, len(list)-1)
	return q.List
}

func (q *quickSort3Way) sort(lo, hi int) {
	if lo < hi {
		p1, p2 := q.partition(lo, hi)
		q.sort(lo, p1-1)
		q.sort(p1+1, p2-1)
		q.sort(p2+1, hi)
	}
}

func (q *quickSort3Way) partition(lo, hi int) (int, int) {
	pivot := q.List[lo]
	lt, gt := lo, hi

	i := lo
	for i <= gt {
		cmp := q.List[i]
		if cmp < pivot {
			q.List[i], q.List[lt] = q.List[lt], q.List[i]
			i++
			lt++
		} else if cmp > pivot {
			q.List[i], q.List[gt] = q.List[gt], q.List[i]
			gt--
		} else {
			i++
		}

	}

	return lt, gt
}
