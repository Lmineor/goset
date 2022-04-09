package goset

import (
	"testing"
)

func TestStringSet(t *testing.T) {
	s := String{}
	s2 := String{}
	if len(s) != 0 {
		t.Errorf("Expected len=0: %d", len(s))
	}
	s.Add("a", "b")
	if len(s) != 2 {
		t.Errorf("Expected len=2: %d", len(s))
	}
	s.Add("c")
	if s.HasItem("d") {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.HasItem("a") {
		t.Errorf("Missing contents: %#v", s)
	}
	s.Remove("a")
	if s.HasItem("a") {
		t.Errorf("Unexpected contents: %#v", s)
	}
	s.Add("a")
	if s.HasAll("a", "b", "d") {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.HasAll("a", "b") {
		t.Errorf("Missing contents: %#v", s)
	}
	s2.Add("a", "b", "d")
	if s.IsSuperset(s2) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	s2.Remove("d")
	if !s.IsSuperset(s2) {
		t.Errorf("Missing contents: %#v", s)
	}
}

func TestIntSet(t *testing.T) {
	i := Int{}
	i2 := Int{}
	if len(i) != 0 {
		t.Errorf("Expected len=0: %d", len(i))
	}
	i.Add(1, 2)
	if len(i) != 2 {
		t.Errorf("Expected len=2: %d", len(i))
	}
	i.Add(3)
	if i.HasItem(4) {
		t.Errorf("Unexpected contents: %#v", i)
	}
	if !i.HasItem(2) {
		t.Errorf("Missing contents: %#v", i)
	}
	i.Remove(1)
	if i.HasItem(1) {
		t.Errorf("Unexpected contents: %#v", i)
	}
	i.Add(1)
	if i.HasAll(1, 4, 5) {
		t.Errorf("Unexpected contents: %#v", i)
	}
	if !i.HasAll(1, 2) {
		t.Errorf("Missing contents: %#v", i)
	}
	i2.Add(1, 2, 5)
	if i.IsSuperset(i2) {
		t.Errorf("Unexpected contents: %#v", i)
	}
	i2.Remove(5)
	if !i.IsSuperset(i2) {
		t.Errorf("Missing contents: %#v", i)
	}
}
