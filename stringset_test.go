package goset

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewStringSetAdd(t *testing.T) {
	ss := NewStringSet()
	ss.Add("s", "a", "b", "c", "s")
	res1 := ss.SortList()
	fmt.Println(res1)
}

func TestAdd(t *testing.T) {
	ans := []string{"a", "b", "c", "s", "ss"}
	stringSet := NewStringSet()
	stringSet.Add("ss", "a", "b", "c", "s", "s", "s", "s")
	res := stringSet.SortList()
	if !equalSlice(ans, res) {
		t.Error("failed to Add")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Tests Begins..")
	m.Run()
}

func BenchmarkAdd(b *testing.B) {
	ss := NewStringSet()
	for i := 0; i < b.N; i++ {
		si := strconv.Itoa(i)
		ss.Add(si)
	}
}

func equalSlice(s1 []string, s2 []string) bool {
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
