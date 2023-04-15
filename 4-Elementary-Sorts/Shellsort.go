package ElementarySorts

// ShellSort N^3/2
func ShellSort[ordered Ordered](list []ordered, callback OrderedCallback[ordered]) {
	lenList := len(list)
	step := 1
	for step < lenList/3 {
		step = 3*step + 1
	}

	for step >= 1 {
		for i := 0; i < lenList; i++ {
			step = 1
			for j := i; j >= step; j -= step {
				if callback(list[j], list[j-step]) {
					list[j], list[j-step] = list[j-step], list[j]
				}
			}
		}

		step = step / 3
	}
}
