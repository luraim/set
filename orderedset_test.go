package set

import (
	"reflect"
	"testing"
)

func TestOrderedSet(t *testing.T) {
	s := NewOrdered(1, 2, 3, 4, 5)
	expectedLst := []int{1, 2, 3, 4, 5}
	gotLst := s.ToList()
	if !reflect.DeepEqual(gotLst, expectedLst) {
		t.Errorf("after creation: expected %v, got %v", expectedLst, gotLst)
	}

	if s.IsEmpty() {
		t.Errorf("empty check failed")
	}

	if s.String() != "(1, 2, 3, 4, 5)" {
		t.Errorf("string repr failed. Got %s", s.String())
	}

	if s.Count() != 5 {
		t.Errorf("count check failed. Expected 5, got: %d", s.Count())
	}

	s.Add(3, 7, 6, 7, 1, 2)
	expected := NewOrdered(1, 2, 3, 4, 5, 7, 6)
	if !expected.IsEqual(s) {
		t.Errorf("after addition: expected %v, got %v", expected, s)
	}

	s.Remove(4)
	expected = NewOrdered(1, 2, 3, 5, 7, 6)
	if !expected.IsEqual(s) {
		t.Errorf("after removal: expected %v, got %v", expected, s)
	}

	s.Remove(100)
	expected = NewOrdered(1, 2, 3, 5, 7, 6)
	if !expected.IsEqual(s) {
		t.Errorf("after fake removal: expected %v, got %v", expected, s)
	}

	if s.Contains(3) == false {
		t.Errorf("positive check for presence failed")
	}

	if s.Contains(4) == true {
		t.Errorf("negative check for presence failed")
	}

	newS := NewOrdered(1, 7, 10, 20, 30)

	expected = NewOrdered(1, 2, 3, 5, 7, 6, 10, 20, 30)
	got := s.Union(newS)
	if !expected.IsEqual(got) {
		t.Errorf("after union: expected %v, got %v", expected, got)
	}

	expected = NewOrdered(1, 7)
	got = s.Intersection(newS)
	if !expected.IsEqual(got) {
		t.Errorf("after intersection: expected %v, got %v", expected, got)
	}

	expected = NewOrdered(2, 3, 5, 6)
	got = s.Diff(newS)
	if !expected.IsEqual(got) {
		t.Errorf("after diff: expected %v, got %v", expected, got)
	}
}
