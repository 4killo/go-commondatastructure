package set

import (
	"sync"
)

//Set data structure
type Set struct {
	set  map[interface{}]bool
	sync sync.RWMutex
}

//New Create new set
func New() *Set {
	return &Set{set: make(map[interface{}]bool, 0)}
}

//Add ...
func (s *Set) Add(item interface{}) bool {
	s.sync.Lock()
	defer s.sync.Unlock()
	if _, ok := s.set[item]; !ok {
		s.set[item] = true
		return true
	}
	return false
}

//Contains ...
func (s *Set) Contains(item interface{}) bool {
	s.sync.Lock()
	defer s.sync.Unlock()
	_, ok := s.set[item]
	return ok
}

//Remove ...
func (s *Set) Remove(item interface{}) {
	s.sync.Lock()
	defer s.sync.Unlock()
	delete(s.set, item)
}

//Size ...
func (s *Set) Size() int {
	s.sync.Lock()
	defer s.sync.Unlock()
	return len(s.set)
}

//AllItem ...
func (s *Set) AllItem() []interface{} {
	s.sync.Lock()
	defer s.sync.Unlock()
	items := make([]interface{}, len(s.set))
	i := 0
	for key := range s.set {
		items[i] = key
		i++
	}

	return items
}
