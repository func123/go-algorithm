package sort

// 自顶向下归并排序
type DownMerge struct {
}

func (s *DownMerge) Sort(s1 []int) bool {
	//isAlreadySorted := true
	aux := make([]int, len(s1))
	s.doSort(&s1, &aux, 0, len(s1)-1)
	return true
}

// 对数组索引区间【lo ~ hi】的值进行排序
func (s *DownMerge) doSort(s1, aux *[]int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	// 利用分治的思想：分成两个子数组，分别排序后合并
	s.doSort(s1, aux, lo, mid)
	s.doSort(s1, aux, mid+1, hi)
	// 将两个 有序 数组合并
	s.merge(s1, aux, lo, mid, hi)
}

// merge 前提是 lo - mid ,mid+1 - hi ，这两个被归并的数组是有序的
// 辅助数组 在merge中充当原数组快照的作用，因此不关心之前的值，这里会在需要的区间【lo~hi】复制一份快照数据
func (s *DownMerge) merge(s1, aux *[]int, lo, mid, hi int) {
	// 设置2个数组的起点：
	// A : lo~mid
	// B : mid+1~hi
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		(*aux)[k] = (*s1)[k]
	}
	for k := lo; k <= hi; k++ {
		if i > mid { // 数组A已经全放在s1中，后面全都放B数组的数据
			(*s1)[k] = (*aux)[j]
			j += 1
		} else if j > hi { // 数组B已经全放在s1中，后面全都放A数组的数据
			(*s1)[k] = (*aux)[i]
			i += 1
		} else if Less((*aux)[j], (*aux)[i]) { // 数组B的j值比A的i值小，把B的j值放到s1中
			(*s1)[k] = (*aux)[j]				// 下一个对比数组B的j+1值和数组A的i值
			j += 1
		} else { // 数组B的j值 比 A 的i值大，把A 的i值放到s1中
			(*s1)[k] = (*aux)[i]
			i += 1
		}
	}
}
