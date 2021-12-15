package stringset

import (
	"sort"
	"sync"
)

type StringSet struct {
	sync.RWMutex
	m map[string]bool
}

// New 新建集合对象
func NewStringSet(items ...string) *StringSet {
	s := &StringSet{m: make(map[string]bool, len(items))}
	s.Add(items...)
	return s
}

// Add 添加元素
func (s *StringSet) Add(items ...string) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		s.m[item] = true
	}
}

// Remove 删除元素
func (s *StringSet) Remove(items ...string) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		delete(s.m, item)
	}
}

// In 判断元素是否在集合中
func (s *StringSet) Contains(items ...string) bool {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		if _, ok := s.m[item]; !ok {
			return false
		}
	}
	return true
}

func (s *StringSet) Count() int {
	return len(s.m)
}

func (s *StringSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]bool{}
}

// 空集合判断
func (s *StringSet) Empty() bool {
	return len(s.m) == 0
}

// List 返回无序列表
func (s *StringSet) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := make([]string, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

// SortList 返回排序列表
func (s *StringSet) SortList() []string {
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
func (s *StringSet) Intersect(sets ...*StringSet) *StringSet {
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

// Minux 返回集合的差集
func (s *StringSet) Minus(sets ...*StringSet) *StringSet {
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
func (s *StringSet) Union(sets ...*StringSet) *StringSet {
	rerfer := NewStringSet(s.List()...)
	for _, set := range sets {
		for e := range set.m {
			rerfer.m[e] = true
		}
	}
	return rerfer
}

// Complement 补集
func (s *StringSet) Complement(full *StringSet) *StringSet {
	rerfer := NewStringSet()
	for e := range full.m {
		if _, ok := s.m[e]; !ok {
			rerfer.Add(e)
		}
	}
	return rerfer
}
