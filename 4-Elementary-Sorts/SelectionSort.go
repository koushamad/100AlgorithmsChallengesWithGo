package ElementarySorts

func SelectionSort[ordered Ordered](list []ordered, callback OrderedCallback[ordered]) {
	lenList := len(list)

	for i := 0; i < lenList; i++ {
		for j := i + 1; j < lenList; j++ {

			if callback(list[i], list[j]) {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
