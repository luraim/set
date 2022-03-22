package set

import (
	"fmt"
	"strings"

	"github.com/luraim/fun"
)

type OrderedSet[T comparable] struct {
	Elems []T        `json:"set_elems"`
	EMap  map[T]bool `json:"set_e_map"`
}

func (s OrderedSet[T]) String() string {
	elemStrs := fun.Map(s.Elems, func(e T) string {
		return fmt.Sprintf("%v", e)
	})
	return fmt.Sprintf("(%s)", strings.Join(elemStrs, ", "))
}

func NewOrdered[T comparable](elems ...T) *OrderedSet[T] {
	ret := &OrderedSet[T]{
		Elems: make([]T, 0),
		EMap:  make(map[T]bool),
	}
	ret.Add(elems...)
	return ret
}

func (s *OrderedSet[T]) IsEmpty() bool {
	return len(s.Elems) == 0
}

func (s *OrderedSet[T]) ToList() []T {
	return s.Elems
}

func (s *OrderedSet[T]) Contains(e T) bool {
	_, ok := s.EMap[e]
	return ok
}

// Count returns the number of elements in the set
func (s *OrderedSet[T]) Count() int {
	return len(s.Elems)
}

// Add one or more elements to the set
func (s *OrderedSet[T]) Add(elems ...T) {
	for _, e := range elems {
		if !s.Contains(e) {
			s.Elems = append(s.Elems, e)
			s.EMap[e] = true
		}
	}
}

// Union returns a new set that has all elements from both sets
func (s *OrderedSet[T]) Union(other *OrderedSet[T]) *OrderedSet[T] {
	ret := NewOrdered[T]()
	ret.Add(s.Elems...)
	ret.Add(other.Elems...)
	return ret
}

// Intersection returns a new set that has common
// elements present in both sets
func (s *OrderedSet[T]) Intersection(other *OrderedSet[T]) *OrderedSet[T] {
	ret := NewOrdered[T]()
	for _, e := range s.Elems {
		if other.Contains(e) {
			ret.Add(e)
		}
	}
	return ret
}

// Equals returns true if every element is same between this and the
// other set, and in the same order
func (s *OrderedSet[T]) IsEqual(other *OrderedSet[T]) bool {
	if len(s.Elems) != len(other.Elems) {
		return false
	}
	// all elements should be the same and also in the same order
	for i := 0; i < len(s.Elems); i++ {
		if s.Elems[i] != other.Elems[i] {
			return false
		}
	}
	return true
}

// Diff returns a new set with elements in this set, but not in the
// other set
func (s *OrderedSet[T]) Diff(other *OrderedSet[T]) *OrderedSet[T] {
	ret := NewOrdered[T]()
	for _, e := range s.Elems {
		if !other.Contains(e) {
			ret.Add(e)
		}
	}
	return ret
}

// Remove the given elements from the current set
func (s *OrderedSet[T]) Remove(elems ...T) {
	for _, e := range elems {
		if s.Contains(e) {
			// remove from list
			for i, v := range s.Elems {
				if v == e {
					// found element, remove it
					s.Elems = append(s.Elems[:i], s.Elems[i+1:]...)
				}
			}
			// remove from map
			delete(s.EMap, e)
		}
	}
}
