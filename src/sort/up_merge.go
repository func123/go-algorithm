package sort

// 自底向上归并排序
type UpMerge struct {
}

func (u *UpMerge) Sort(s1 []int) bool {
	//isAlreadySorted := true
	aux := make([]int, len(s1))
	u.doSort(&s1, &aux, 0, len(s1)-1)
	return true
}

func (u *UpMerge) doSort(s1, aux *[]int, lo, hi int) {
	mid := lo + (hi-lo)/2
	for i := 1; i <= mid; i <<= 1 {
		for k := 0; k < hi-i; k += i + i {
			u.merge(s1, aux, k, k+k-1, u.min(hi, k+k-1))
		}
	}

}

// merge 前提是 lo - mid ,mid+1 - hi ，这两个被归并的数组是有序的
// 辅助数组 在merge中充当原数组快照的作用，因此不关心之前的值，这里会在需要的区间【lo~hi】复制一份快照数据
func (u *UpMerge) merge(s1, aux *[]int, lo, mid, hi int) {
	// 设置2个数组的起点：
	// A :  lo~mid
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
			(*s1)[k] = (*aux)[j] // 下一个对比数组B的j+1值和数组A的i值
			j += 1
		} else { // 数组B的j值 比 A 的i值大，把A 的i值放到s1中
			(*s1)[k] = (*aux)[i]
			i += 1
		}
	}
}

func (u *UpMerge) min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
