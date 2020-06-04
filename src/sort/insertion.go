package sort

// 插入排序
type Insertion struct {
}

// 第一次循环i: 用于定位数组中的排序进度点
// 第二次循环j：因为要把当前i的值放到数组进度点左边合适的位置，因此这个i的index只是提供一个定位起点的作用，绝对不能在这个循环中作为变量被修改
func (s *Insertion) Sort(s1 []int) bool {
	isAlreadySorted := true
	arrayLen := len(s1)
	for i := 1; i < arrayLen; i++ {
		for j := i - 1; j >= 0 && Less(s1[j+1], s1[j]); j-- {
			//if Less(s1[i], s1[j]) { 这里不能用i
			Exch(s1, j+1, j)
			isAlreadySorted = false
		}
		continue
	}
	return isAlreadySorted
}
