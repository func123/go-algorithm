package sort

import (
	"math/rand"
)

type Sorter interface {
	Sort(s []int) bool
}

func Less(s1, s2 int) bool {
	if s1 >= s2 {
		return false
	}
	return true
}

func IsSorted(s1 []int) bool {
	for i := len(s1) - 1; i > 0; i-- {
		if s1[i] < s1[i-1] {
			return false
		}
	}
	return true
}

func Exch(s []int, s1, s2 int) {
	temp := s[s1]
	s[s1] = s[s2]
	s[s2] = temp
}

func GetRandomArrays(num int) []int {
	s1 := rand.Perm(num)
	return s1
}
