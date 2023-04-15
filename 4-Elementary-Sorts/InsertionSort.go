package ElementarySorts

// InsertionSort ~1/4 N^2
func InsertionSort[ordered Ordered](list []ordered, callback OrderedCallback[ordered]) {
	lenList := len(list)

	for i := 0; i < lenList; i++ {
		for j := i; j > 0; j-- {
			if callback(list[j], list[j-1]) {
				list[j], list[j-1] = list[j-1], list[j]
			} else {
				break
			}
		}
	}
}
