package main

import (
	"sync"
)

// Set type of set with Mutex
type Set struct {
	sync.Mutex
	mm map[int]struct{}
}

// NewSet return new set
func NewSet() *Set {
	return &Set{
		mm: map[int]struct{}{},
	}
}

// Add add new element to set
func (s *Set) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

// Has read element from set
func (s *Set) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

// SetRW type of set with RWMutex
type SetRW struct {
	sync.RWMutex
	mm map[int]struct{}
}

// NewSetRW return new setRW
func NewSetRW() *SetRW {
	return &SetRW{
		mm: map[int]struct{}{},
	}
}

// Add add new element to set
func (s *SetRW) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

// Has read element from set
func (s *SetRW) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}
