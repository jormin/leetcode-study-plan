package day09

import (
	"testing"
)

// TestMyQueue 测试用栈实现队列
func TestMyQueue(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	v := obj.Peek()
	if v != 1 {
		t.Errorf("Peek() = %v, want %v", v, 1)
	}
	v = obj.Pop()
	if v != 1 {
		t.Errorf("Pop() = %v, want %v", v, 1)
	}
	b := obj.Empty()
	if b != false {
		t.Errorf("Empty() = %v, want %v", b, false)
	}
	v = obj.Pop()
	if v != 2 {
		t.Errorf("Pop() = %v, want %v", v, 2)
	}
	obj.Push(3)
	v = obj.Pop()
	if v != 3 {
		t.Errorf("Pop() = %v, want %v", v, 3)
	}
}
