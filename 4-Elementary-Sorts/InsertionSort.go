package ElementarySorts

// InsertionSort ~1/4 N^2
func InsertionSort[ordered Ordered](list []ordered, callback OrderedCallback[ordered]) {
	lenList := len(list)

	for i := 0; i < lenList; i++ {
		for j := 0; j <= i; j++ {
			if callback(list[i], list[j]) {
				list[i], list[j] = list[j], list[i]
			} else {
				break
			}
		}
	}
}
