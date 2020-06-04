package sort

type Selection struct {
}

func (s *Selection) Sort(s1 []int) bool {
	isAlreadySorted := true
	arrayLen := len(s1)
	for i := 0; i < arrayLen; i++ {
		min := i
		for j := i + 1; j < arrayLen; j++ {
			if Less(s1[j], s1[min]) {
				isAlreadySorted = false
				min = j
			}
		}
		Exch(s1, i, min)
	}
	return isAlreadySorted
}

