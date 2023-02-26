package goset

import (
	"sort"
)

type Int map[int]Empty

// NewInt 新建集合对象
func NewInt(items ...int) Int {
	is := Int{}
	is.Add(items...)
	return is
}

// Add 添加元素
func (i Int) Add(items ...int) Int {
	for _, item := range items {
		i[item] = Empty{}
	}
	return i
}

// Remove 删除元素
func (i Int) Remove(items ...int) Int {
	for _, item := range items {
		delete(i, item)
	}
	return i
}

// HasItem 判断元素是否在集合中
func (i Int) HasItem(item int) bool {
	_, has := i[item]
	return has
}

// HasAny 只有有元素在集合中就行
func (i Int) HasAny(items ...int) bool {
	for _, item := range items {
		if i.HasItem(item) {
			return true
		}
	}
	return false
}

// HasAll items 全部都在集合中
func (i Int) HasAll(items ...int) bool {
	for _, item := range items {
		if !i.HasItem(item) {
			return false
		}
	}
	return true
}

// Count 当前集合的长度
func (i Int) Count() int {
	return len(i)
}

func (i Int) Clear() {
	allItems := i.List()
	i.Remove(allItems...)
}

// Empty 空集合判断
func (i Int) Empty() bool {
	return i.Count() == 0
}

// List 返回无序列表
func (i Int) List() []int {
	list := make([]int, 0, i.Count())
	for item := range i {
		list = append(list, item)
	}
	return list
}

// SortedList 返回排序列表
func (i Int) SortedList() []int {
	list := make([]int, 0, len(i))
	for item := range i {
		list = append(list, item)
	}
	sort.Ints(list)
	return list
}

// Difference returns a set of objects that are not in s2
func (i Int) Difference(i2 Int) Int {
	refer := NewInt()
	for key := range i {
		if !i2.HasItem(key) {
			refer.Add(key)
		}
	}
	return refer
}

// Intersection 返回集合的交集
func (i Int) Intersection(i2 Int) Int {
	var longer, shorter Int
	refer := NewInt()
	if i.Count() < i2.Count() {
		longer = i2
		shorter = i
	} else {
		longer = i
		shorter = i2
	}
	for item := range shorter {
		if longer.HasItem(item) {
			refer.Add(item)
		}
	}
	return refer
}

// Union 返回集合的并集
func (i Int) Union(i2 Int) Int {
	refer := NewInt()
	for key := range i {
		refer.Add(key)
	}
	for key := range i2 {
		refer.Add(key)
	}
	return refer
}

func (i Int) IsSuperset(i2 Int) bool {
	for item := range i2 {
		if !i.HasItem(item) {
			return false
		}
	}
	return true
}

func (i Int) Equal(i2 Int) bool {
	return i.Count() == i2.Count() && i.IsSuperset(i2)
}
