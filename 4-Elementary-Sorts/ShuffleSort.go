package ElementarySorts

import (
	"math/rand"
	"time"
)

func ShuffleSort[ordered Ordered](list []ordered) {
	lenList := len(list)

	for i := 0; i < lenList; i++ {
		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		r1.Shuffle(lenList, func(i, j int) {
			list[i], list[j] = list[j], list[i]
		})
	}
}
