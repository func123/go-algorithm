package sort

type Shell struct {
}

func (s *Shell) Sort(s1 []int) bool {
	isAlreadySorted := true
	arrayLen := len(s1)
	h := 1
	for ; h < arrayLen/3; h = h*3 + 1 {}
	for ; h >= 1; h = h / 3 { // 设置 比较跨度
		for i := h; i < arrayLen; i++ { // 定位此次比较元素的位置，跟数组中该小标左边的元素比较，此次比较跨度为h
			// 例如h=4 : i=4,i=5,i=6,i=7,i=8(此次就跟i=4的下标比较)`````i=12（此次跟8 和 4 比较）
			for j := i; j >= h && Less(s1[j], s1[j-h]); j -= h { // 向左开始逐步比较
				Exch(s1, j, j-h)
				isAlreadySorted = false
			}
		}
	}
	return isAlreadySorted
}
