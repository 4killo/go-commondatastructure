package stack

import "sync"

//TODO: allow generic type of  stack creation like Stack<int>.

// Stack holds items
type Stack struct {
	item []interface{}
	sync sync.RWMutex
}

//New creates new stack
func New() *Stack {
	return &Stack{item: make([]interface{}, 0)}
}

//Push ...
func (s *Stack) Push(i interface{}) {
	s.sync.Lock()
	defer s.sync.Unlock()
	s.item = append(s.item, i)

}

//Pop ...
func (s *Stack) Pop() interface{} {
	s.sync.Lock()
	defer s.sync.Unlock()
	i := s.item[len(s.item)-1]
	s.item = s.item[0 : len(s.item)-1]
	return i
}

//Size of the stack
func (s *Stack) Size() int {
	s.sync.Lock()
	defer s.sync.Unlock()
	l := len(s.item)
	return l
}
