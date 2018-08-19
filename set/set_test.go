package set

import (
	"fmt"
	"testing"
)

var s = New()

func TestAdd(t *testing.T) {
	s.Add(1)
	s.Add(2)
	s.Add(1)
	s.Add(3)
	s.Add(3)

	if s.Size() != 3 {
		t.Errorf("Wrong size expected %v got %v", 3, s.Size())
	}

}

func TestContains(t *testing.T) {
	if !s.Contains(1) {
		t.Errorf("Item expected")
	}
	if !s.Contains(2) {
		t.Errorf("Item expected")
	}
	if !s.Contains(3) {
		t.Errorf("Item expected")
	}
	if s.Contains(4) {
		t.Errorf("Item not expected")
	}
}

func TestRemove(t *testing.T) {
	s.Remove(1)
	if s.Contains(1) {
		t.Errorf("Item expected")
	}
}

func TestAllItem(t *testing.T) {
	items := s.AllItem()
	for _, val := range items {
		fmt.Println(val)
	}
}
