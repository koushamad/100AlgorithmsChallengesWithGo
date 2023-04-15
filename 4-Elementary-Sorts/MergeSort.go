package ElementarySorts

// MergeSort NlogN
type mergeSort struct {
	List []int
	tmp  []int
}

func MergeSort(list []int) []int {
	m := &mergeSort{
		List: list,
		tmp:  make([]int, len(list)),
	}
	m.sort(0, len(list)-1)
	return m.List
}

func (m *mergeSort) sort(low, hig int) {
	if hig <= low {
		return
	}

	mid := low + (hig-low)/2
	m.sort(low, mid)
	m.sort(mid+1, hig)
	m.merge(low, mid, hig)
}

func (m *mergeSort) merge(low, mid, hig int) {
	for k := low; k <= hig; k++ {
		m.tmp[k] = m.List[k]
	}

	i, j := low, mid+1

	for k := low; k <= hig; k++ {
		if i > mid {
			m.List[k] = m.tmp[j]
			j++
		} else if j > hig {
			m.List[k] = m.tmp[i]
			i++
		} else if m.tmp[j] < m.tmp[i] {
			m.List[k] = m.tmp[j]
			j++
		} else {
			m.List[k] = m.tmp[i]
			i++
		}
	}
}
