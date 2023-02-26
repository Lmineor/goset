package goset

import (
	"sort"
)

type String map[string]Empty

// NewString 新建集合对象
func NewString(items ...string) String {
	ss := String{}
	ss.Add(items...)
	return ss
}

// Add 添加元素
func (s String) Add(items ...string) String {
	for _, item := range items {
		s[item] = Empty{}
	}
	return s
}

// Remove 删除元素
func (s String) Remove(items ...string) String {
	for _, item := range items {
		delete(s, item)
	}
	return s
}

// HasItem 判断元素是否在集合中
func (s String) HasItem(item string) bool {
	_, has := s[item]
	return has
}

// HasAny 只有有元素在集合中就行
func (s String) HasAny(items ...string) bool {
	for _, item := range items {
		if s.HasItem(item) {
			return true
		}
	}
	return false
}

// HasAll items 全部都在集合中
func (s String) HasAll(items ...string) bool {
	for _, item := range items {
		if !s.HasItem(item) {
			return false
		}
	}
	return true
}

// Count 当前集合的长度
func (s String) Count() int {
	return len(s)
}

func (s String) Clear() {
	allItems := s.List()
	s.Remove(allItems...)
}

// Empty 空集合判断
func (s String) Empty() bool {
	return s.Count() == 0
}

// List 返回无序列表
func (s String) List() []string {
	list := make([]string, 0, s.Count())
	for item := range s {
		list = append(list, item)
	}
	return list
}

// SortedList 返回排序列表
func (s String) SortedList() []string {
	list := make([]string, 0, len(s))
	for item := range s {
		list = append(list, item)
	}
	sort.Strings(list)
	return list
}

// Difference returns a set of objects that are not in s2
func (s String) Difference(s2 String) String {
	refer := NewString()
	for key := range s {
		if !s2.HasItem(key) {
			refer.Add(key)
		}
	}
	return refer
}

// Intersection 返回集合的交集
func (s String) Intersection(s2 String) String {
	var longer, shorter String
	refer := NewString()
	if s.Count() < s2.Count() {
		longer = s2
		shorter = s
	} else {
		longer = s
		shorter = s2
	}
	for item := range shorter {
		if longer.HasItem(item) {
			refer.Add(item)
		}
	}
	return refer
}

// Union 返回集合的并集
func (s String) Union(s2 String) String {
	refer := NewString()
	for key := range s {
		refer.Add(key)
	}
	for key := range s2 {
		refer.Add(key)
	}
	return refer
}

func (s String) IsSuperset(s2 String) bool {
	for item := range s2 {
		if !s.HasItem(item) {
			return false
		}
	}
	return true
}

func (s String) Equal(s2 String) bool {
	return s.Count() == s2.Count() && s.IsSuperset(s2)
}
