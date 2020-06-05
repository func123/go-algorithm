package sort

import (
	"testing"
	"time"
)

func doSortTesting(sorter Sorter, s1 []int, t *testing.T) {
	start := time.Now()
	alreadySorted := sorter.Sort(s1)
	if IsSorted(s1) {
		t.Logf("sort successfully,already sorted(%v),spend time(%v)", alreadySorted, time.Since(start))
	} else {
		t.Logf("sort unsuccessfully,spend time(%v)", time.Since(start))
	}
	t.Log(s1)
}

// 选择排序
//func TestSelection(t *testing.T) {
//	s1 := GetRandomArrays(100000)
//	sorter := new(Selection)
//	doSortTesting(sorter, s1, t)
//}

// 插入排序
//func TestInsertion(t *testing.T) {
//	s1 := GetRandomArrays(100000)
//	sorter := new(Insertion)
//	doSortTesting(sorter, s1, t)
//}

// 希尔排序
//func TestShell(t *testing.T) {
//	s1 := GetRandomArrays(100000)
//	sorter := new(Shell)
//	doSortTesting(sorter, s1, t)
//}

// 自顶向下归并排序
//func TestDownMerge(t *testing.T) {
//	s1 := GetRandomArrays(100000)
//	sorter := new(DownMerge)
//	doSortTesting(sorter, s1, t)
//}

// 自底向上归并排序
func TestUpMerge(t *testing.T) {
	s1 := GetRandomArrays(10)
	sorter := new(UpMerge)
	doSortTesting(sorter, s1, t)
}

