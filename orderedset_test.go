package set

import (
	"reflect"
	"testing"
)

func TestOrderedSet(t *testing.T) {
	s := NewOrdered(1, 2, 3, 4, 5)
	expected := []int{1, 2, 3, 4, 5}
	got := s.ToList()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("after creation: expected %v, got %v", expected, got)
	}

	s.Add(3, 7, 6, 7, 1, 2)
	expected = []int{1, 2, 3, 4, 5, 7, 6}
	got = s.ToList()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("after addition: expected %v, got %v", expected, got)
	}

	s.Remove(4)
	expected = []int{1, 2, 3, 5, 7, 6}
	got = s.ToList()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("after removal: expected %v, got %v", expected, got)
	}

	s.Remove(100)
	expected = []int{1, 2, 3, 5, 7, 6}
	got = s.ToList()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("after fake removal: expected %v, got %v", expected, got)
	}

	if s.Contains(3) == false {
		t.Errorf("positive check for presence failed")
	}

	if s.Contains(4) == true {
		t.Errorf("negative check for presence failed")
	}

	newS := NewOrdered(1, 7, 10, 20, 30)

	expected = []int{1, 2, 3, 5, 7, 6, 10, 20, 30}
	got = s.Union(newS).ToList()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("after union: expected %v, got %v", expected, got)
	}

	expected = []int{1, 7}
	got = s.Intersection(newS).ToList()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("after union: expected %v, got %v", expected, got)
	}

}
