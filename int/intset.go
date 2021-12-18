package int

import (
	"sort"
	"sync"
)

type intSet struct {
	sync.RWMutex
	m map[int]bool
}

// NewIntSet 新建集合对象
func NewIntSet(items ...int) *intSet {
	s := &intSet{m: make(map[int]bool, len(items))}
	s.Add(items...)
	return s
}

// Add 添加元素
func (s *intSet) Add(items ...int) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		s.m[item] = true
	}
}

// Remove 删除元素
func (s *intSet) Remove(items ...int) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		delete(s.m, item)
	}
}

// Contains 判断元素是否在集合中
func (s *intSet) Contains(items ...int) bool {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		if _, ok := s.m[item]; !ok {
			return false
		}
	}
	return true
}

func (s *intSet) Count() int {
	return len(s.m)
}

func (s *intSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
}

// Empty 空集合判断
func (s *intSet) Empty() bool {
	return len(s.m) == 0
}

// List 返回无序列表
func (s *intSet) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

// SortList 返回排序列表
func (s *intSet) SortList() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	sort.Ints(list)
	return list
}

// Intersect 返回集合的交集
func (s *intSet) Intersect(sets ...*intSet) *intSet {
	refer := NewIntSet(s.List()...)
	for _, set := range sets {
		for e := range s.m {
			if _, ok := set.m[e]; !ok {
				delete(refer.m, e)
			}
		}
	}
	return refer
}

// Minus 返回集合的差集
func (s *intSet) Minus(sets ...*intSet) *intSet {
	refer := s.Union(sets...)
	for _, set := range sets {
		for e := range set.m {
			if _, ok := s.m[e]; ok {
				delete(refer.m, e)
			}
		}
	}
	return refer
}

// Union 返回集合的并集
func (s *intSet) Union(sets ...*intSet) *intSet {
	refer := NewIntSet(s.List()...)
	for _, set := range sets {
		for e := range set.m {
			refer.m[e] = true
		}
	}
	return refer
}

// Complement 补集
func (s *intSet) Complement(full *intSet) *intSet {
	refer := NewIntSet()
	for e := range full.m {
		if _, ok := s.m[e]; !ok {
			refer.Add(e)
		}
	}
	return refer
}
