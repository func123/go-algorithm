package sort

// 自顶向下归并排序
type DownMerge struct {
}

func (s *DownMerge) Sort(s1 []int) bool {
	//isAlreadySorted := true
	aux := make([]int,  len(s1))
	copy(aux, s1)
	s.doSort(&s1, &aux, 0, len(s1)-1)
	return true
}

func (s *DownMerge) doSort(s1, aux *[]int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	s.doSort(s1, aux, lo, mid)
	s.doSort(s1, aux, mid+1, hi)
	s.merge(s1, aux, lo, mid, hi)
}

// merge 前提是 lo - mid ,mid+1 - hi ，这两个被归并的数组是有序的
func (s *DownMerge) merge(s1, aux *[]int, lo, mid, hi int) {
	i, j := lo, mid+1
	for k := lo; k < hi; k++ {
		(*aux)[k] = (*s1)[k]
	}
	for k := lo; k < hi; k++ {
		if i > mid {
			(*s1)[k] = (*aux)[j]
			j += 1
		} else if j > hi {
			(*s1)[k] = (*aux)[i]
			i += 1
		} else if Less((*aux)[j], (*aux)[i]) {
			(*s1)[k] = (*aux)[j]
			j += 1
		} else {
			(*s1)[k] = (*aux)[i]
			i += 1
		}
	}
}
