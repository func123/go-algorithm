package sort

// 自底向上归并排序
type UpMerge struct {
}

func (u *UpMerge) Sort(s1 []int) bool {
	//isAlreadySorted := true
	length := len(s1)
	aux := make([]int, length)

	// i 可以理解成比较子数组的长度，整个数组长度是length，那么最大子数组的长度是length-1，它对应右比较子数组长度是1
	// for i := 1; i < length-1; i <<= 1（X），因此是用 i < length
	for i := 1; i < length; i <<= 1 {
		//  k 是每组比较数组的起始index
		// 例如每个比较子数组长度为 i = 1，第一组比较数组为【0】和[1]，那么第二组比较数组的起始index = 0 + 1 + 1 = 2 (每组有两个比较数组i+i）····
		// 那么存在在最后一组比较数组中，左数组大小为i ,右数组大小为1时，此时有最大比较起始index，此时的index = length-i-1,
		// 因此 k < length - i
		for k := 0; k < length-i; k += i + i {
			//u.merge(&s1, &aux, k, k-1, u.min(length-1, k+k-1))
			// 因为包含 s1[k] ，因此从起始index = k，开始的子数组长度为i的范围是【k ~ k+i-1】
			u.merge(&s1, &aux, k, k+i-1, u.min(length-1, k+i+i-1))
		}
	}
	return true
}

// merge 前提是 lo - mid ,mid+1 - hi ，这两个被归并的数组是有序的
// 辅助数组 在merge中充当原数组快照的作用，因此不关心之前的值，这里会在需要的区间【lo~hi】复制一份快照数据
func (u *UpMerge) merge(s1, aux *[]int, lo, mid, hi int) {
	// 判断2个数组是否已经有序
	if (*s1)[mid] < (*s1)[mid+1]{
		return
	}
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
