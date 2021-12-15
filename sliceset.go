package goset

import (
	"sort"
	"sync"
)

type SliceSet struct {
	sync.RWMutex
	m map[string]bool
}

// New 新建集合对象
func NewSliceSet(items ...string) *SliceSet {
	s := &SliceSet{m: make(map[string]bool, len(items))}
	s.Add(items...)
	return s
}

// Add 添加元素
func (s *SliceSet) Add(items ...string) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		s.m[item] = true
	}
}

// Remove 删除元素
func (s *SliceSet) Remove(items ...string) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		delete(s.m, item)
	}
}

// In 判断元素是否在集合中
func (s *SliceSet) HasItem(items ...string) bool {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		if _, ok := s.m[item]; !ok {
			return false
		}
	}
	return true
}

func (s *SliceSet) Count() int {
	return len(s.m)
}

func (s *SliceSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]bool{}
}

// 空集合判断
func (s *SliceSet) Empty() bool {
	return len(s.m) == 0
}

// 无序列表
func (s *SliceSet) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := make([]string, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

// 排序列表
func (s *SliceSet) SortList() []string {
	s.RLock()
	defer s.RUnlock()
	list := make([]string, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	sort.Strings(list)
	return list
}

// Intersect is to get the common items between slices
func (s *SliceSet) Intersect(sets ...*SliceSet) *SliceSet {
	refer := NewSliceSet(s.List()...)
	for _, set := range sets {
		for e := range s.m {
			if _, ok := set.m[e]; !ok {
				delete(refer.m, e)
			}
		}
	}
	return refer
}

// Minux is to get the different items between slices
func (s *SliceSet) Minus(sets ...*SliceSet) *SliceSet {
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

// Union is to get the all items between slices
func (s *SliceSet) Union(sets ...*SliceSet) *SliceSet {
	rerfer := NewSliceSet(s.List()...)
	for _, set := range sets {
		for e := range set.m {
			rerfer.m[e] = true
		}
	}
	return rerfer
}

// 补集
func (s *SliceSet) Complement(full *SliceSet) *SliceSet {
	rerfer := NewSliceSet()
	for e := range full.m {
		if _, ok := s.m[e]; !ok {
			rerfer.Add(e)
		}
	}
	return rerfer
}
