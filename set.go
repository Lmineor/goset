package goset

type Set interface {
	Add(items ...interface{})
	Remove(items ...interface{})
	HasItem(items ...interface{}) bool
	Count()
	Clear()
	Empty()
	List() []interface{}
	SortList() []interface{}
	Intersect(sets ...*interface{}) *interface{}
	Minus(sets ...*interface{}) *interface{}
	Union(sets ...*interface{}) *interface{}
	Complement(full *interface{}) *interface{}
}
