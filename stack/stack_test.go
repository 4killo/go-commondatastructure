package stack

import (
	"testing"
)

var s = New()

func TestPush(t *testing.T) {
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("Wrong size expected %v got %v", 3, s.Size())
	}
}

func TestPop(t *testing.T) {
	s.Pop()
	if s.Size() != 2 {
		t.Errorf("Wrong size expected %v got %v", 2, s.Size())
	}
	s.Pop()
	if s.Size() != 1 {
		t.Errorf("Wrong size expected %v got %v", 1, s.Size())
	}
	s.Pop()
	if s.Size() != 0 {
		t.Errorf("Wrong size expected %v got %v", 0, s.Size())
	}

}
