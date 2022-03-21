package set

import (
	"fmt"
	"strings"

	"github.com/luraim/fun"
	"golang.org/x/exp/maps"
)

type Set[T comparable] struct {
	EMap map[T]bool `json:"set_e_map"`
}

func (s Set[T]) String() string {
	elemStrs := fun.Map(s.ToList(), func(e T) string {
		return fmt.Sprintf("%v", e)
	})
	return fmt.Sprintf("(%s)", strings.Join(elemStrs, ","))
}

func New[T comparable](elems ...T) *Set[T] {
	ret := &Set[T]{
		EMap: make(map[T]bool),
	}
	ret.Add(elems...)
	return ret
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.EMap) == 0
}

func (s *Set[T]) ToList() []T {
	return maps.Keys(s.EMap)
}

func (s *Set[T]) Contains(e T) bool {
	_, ok := s.EMap[e]
	return ok
}

// Count returns the number of elements in the set
func (s *Set[T]) Count() int {
	return len(s.EMap)
}

// Add one or more elements to the set
func (s *Set[T]) Add(elems ...T) {
	for _, e := range elems {
		if !s.Contains(e) {
			s.EMap[e] = true
		}
	}
}

// Union returns a new set that has all elements from both sets
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	ret := New[T]()
	ret.Add(s.ToList()...)
	ret.Add(other.ToList()...)
	return ret
}

// Intersection returns a new set that has common
// elements present in both sets
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	ret := New[T]()
	for _, e := range s.ToList() {
		if other.Contains(e) {
			ret.Add(e)
		}
	}
	return ret
}

// Equals returns true if every element is same between this and the
// other set
func (s *Set[T]) Equal(other *Set[T]) bool {
	if s.Count() != other.Count() {
		return false
	}
	return maps.Equal(s.EMap, other.EMap)
}

// Diff returns a new set with elements in this set, but not in the
// other set
func (s *Set[T]) Diff(other *Set[T]) *Set[T] {
	ret := New[T]()
	for _, e := range s.ToList() {
		if !other.Contains(e) {
			ret.Add(e)
		}
	}
	return ret
}

// Remove the given elements from the current set
func (s *Set[T]) Remove(elems ...T) {
	for _, e := range elems {
		if s.Contains(e) {
			// remove from map
			delete(s.EMap, e)
		}
	}
}
