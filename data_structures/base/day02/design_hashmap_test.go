package day02

import (
	"testing"
)

// TestMyHashMap 测试设计哈希映射
func TestMyHashMap(t *testing.T) {
	m := Constructor()
	m.Put(1, 1)
	m.Put(2, 2)
	got := m.Get(1)
	if got != 1 {
		t.Errorf("Get() = %v, want %v", got, 1)
	}
	got = m.Get(3)
	if got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}
	m.Put(2, 1)
	got = m.Get(2)
	if got != 1 {
		t.Errorf("Get() = %v, want %v", got, 1)
	}
	m.Remove(2)
	got = m.Get(2)
	if got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}

}
