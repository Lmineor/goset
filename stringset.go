package goset

import (
	"sort"
	"sync"
)

type stringSet struct {
	sync.RWMutex
	m map[string]bool
}

// NewStringSet 新建集合对象
func NewStringSet(items ...string) *stringSet {
	s := &stringSet{m: make(map[string]bool, len(items))}
	s.Add(items...)
	return s
}

// Add 添加元素
func (s *stringSet) Add(items ...string) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		s.m[item] = true
	}
}

// Remove 删除元素
func (s *stringSet) Remove(items ...string) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		delete(s.m, item)
	}
}

// Contains 判断元素是否在集合中
func (s *stringSet) Contains(items ...string) bool {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		if _, ok := s.m[item]; !ok {
			return false
		}
	}
	return true
}

func (s *stringSet) Count() int {
	return len(s.m)
}

func (s *stringSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]bool{}
}

// Empty 空集合判断
func (s *stringSet) Empty() bool {
	return len(s.m) == 0
}

// List 返回无序列表
func (s *stringSet) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := make([]string, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

// SortList 返回排序列表
func (s *stringSet) SortList() []string {
	s.RLock()
	defer s.RUnlock()
	list := make([]string, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	sort.Strings(list)
	return list
}

// Intersect 返回集合的交集
func (s *stringSet) Intersect(sets ...*stringSet) *stringSet {
	refer := NewStringSet(s.List()...)
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
func (s *stringSet) Minus(sets ...*stringSet) *stringSet {
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
func (s *stringSet) Union(sets ...*stringSet) *stringSet {
	refer := NewStringSet(s.List()...)
	for _, set := range sets {
		for e := range set.m {
			refer.m[e] = true
		}
	}
	return refer
}

// Complement 补集
func (s *stringSet) Complement(full *stringSet) *stringSet {
	refer := NewStringSet()
	for e := range full.m {
		if _, ok := s.m[e]; !ok {
			refer.Add(e)
		}
	}
	return refer
}
