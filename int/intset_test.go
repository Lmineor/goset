package int

import (
	"fmt"
	"testing"
)

func TestNewIntSetAdd(t *testing.T) {
	ss := NewIntSet()
	ss.Add(12, 3, 4, 5, 6, 7)
	res1 := ss.SortList()
	fmt.Println(res1)
}

func TestAdd(t *testing.T) {
	ans := []int{3, 4, 5, 6}
	intSet := NewIntSet()
	intSet.Add(3, 4, 5, 3, 4, 5, 6, 4, 3, 3)
	res := intSet.SortList()
	if !equalIntSlice(ans, res) {
		t.Error("failed to Add")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Tests Begins..")
	m.Run()
}

func equalIntSlice(s1 []int, s2 []int) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	if s1Len != s2Len {
		return false
	}
	for i := 0; i < s1Len; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
