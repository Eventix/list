package list

import (
	"testing"
)

func TestFirstAppend(t *testing.T) {
	l := New()

	a := l.Append("a")

	if l.Len() != l.len || l.len != 1 {
		t.Errorf("List size is wrong, expected 1, got %d (l.Len()); %d (l.len)", l.Len(), l.len)
	}

	if e := l.Front(); e != a {
		t.Error("Unexpected value for front", e)
	}

	if e := l.End(); e != a {
		t.Error("Unexpected value for front", e)

	}

	if e := l.Middle(); e != a {
		t.Error("Unexpected value for front", e)
	}
}

func TestAppendAfter(t *testing.T) {
	l := New()

	a := l.Append("a")
	l.Append("c")
	l.AppendAfter("b", a)

	// This should be ignored
	l.AppendAfter("d", nil)

	testListOrder(t, []string{"a", "b", "c"}, l)
}

func TestFirstPrepend(t *testing.T) {
	l := New()

	a := l.Prepend("a")

	if l.Len() != l.len || l.len != 1 {
		t.Errorf("List size is wrong, expected 1, got %d (l.Len()); %d (l.len)", l.Len(), l.len)
	}

	if e := l.Front(); e != a {
		t.Error("Unexpected value for front", e)
	}

	if e := l.End(); e != a {
		t.Error("Unexpected value for front", e)

	}

	if e := l.Middle(); e != a {
		t.Error("Unexpected value for front", e)
	}
}

func TestPrependBefore(t *testing.T) {
	l := New()

	c := l.Prepend("c")
	l.Prepend("a")
	l.PrependBefore("b", c)

	// This should be ignored
	l.PrependBefore("d", nil)

	testListOrder(t, []string{"a", "b", "c"}, l)
}

func TestAppendChain(t *testing.T) {
	l := New()

	expected := []string{"a", "b", "c"}
	for _, v := range expected {
		l.Append(v)
	}

	testListOrder(t, []string{"a", "b", "c"}, l)
}

func TestPrependChain(t *testing.T) {
	l := New()

	expected := []string{"a", "b", "c"}
	for _, v := range expected {
		l.Prepend(v)
	}

	if l.Middle().Value != "a" {
		t.Errorf("Wrong middle value selected %v\n", l.Middle().Value)
	}

	testListOrder(t, []string{"c", "b", "a"}, l)
}

func TestInterleaving(t *testing.T) {
	l := New()

	l.Append("b")
	l.Prepend("a")
	l.Append("c")
	l.Append("d")

	testListOrder(t, []string{"a", "b", "c", "d"}, l)
}

func TestListAppend(t *testing.T) {
	la, lb := New(), New()

	la.Append("a")
	la.Append("b")

	testListOrder(t, []string{"a", "b"}, la.AppendList(la))

	lb.Append("c")
	lb.Append("d")
	lb.Append("e")

	l := la.AppendList(lb)

	testListOrder(t, []string{"a", "b", "c", "d", "e"}, l)

	// Check empty lists
	la, lb = New(), New()
	la.Append("a")
	la.Append("b")
	la.Append("c")

	l = la.AppendList(lb)
	testListOrder(t, []string{"a", "b", "c"}, l)

	// Reverse order
	l = lb.AppendList(la)
	testListOrder(t, []string{"a", "b", "c"}, l)
}

func TestListPrepend(t *testing.T) {
	la, lb := New(), New()

	la.Append("a")
	la.Append("b")

	testListOrder(t, []string{"a", "b"}, la.PrependList(la))

	lb.Append("c")
	lb.Append("d")
	lb.Append("e")

	l := lb.PrependList(la)

	testListOrder(t, []string{"a", "b", "c", "d", "e"}, l)

	// Check empty lists
	la, lb = New(), New()
	la.Append("a")
	la.Append("b")
	la.Append("c")

	l = la.PrependList(lb)
	testListOrder(t, []string{"a", "b", "c"}, l)

	// Reverse order
	l = lb.PrependList(la)
	testListOrder(t, []string{"a", "b", "c"}, l)
}


func TestListAppendMiddle(t *testing.T) {
	la, lb := New(), New()

	la.Append("a")
	la.Append("e")
	lb.Append("b")
	lb.Append("c")
	lb.Append("d")

	l := la.AppendListMiddle(lb)

	testListOrder(t, []string{"a", "b", "c", "d", "e"}, l)

	// Test with list of size 1
	la, lb = New(), New()

	la.Append("b")
	la.Append("d")
	la.Append("e")
	la.Prepend("a")
	lb.Append("c")

	l = la.AppendListMiddle(lb)
	testListOrder(t, []string{"a", "b", "c", "d", "e"}, l)

	// Test empty lists
	la, lb = New(), New()
	la.Append("a")
	la.Append("b")
	la.Append("c")

	l = la.AppendListMiddle(lb)
	testListOrder(t, []string{"a", "b", "c"}, l)

	l = lb.AppendListMiddle(la)
	testListOrder(t, []string{"a", "b", "c"}, l)

	// With middle at the end
	la, lb = New(), New()

	la.Prepend("b")
	la.Prepend("a")
	lb.Append("c")
	lb.Append("d")

	l = la.AppendListMiddle(lb)
	testListOrder(t, []string{"a", "b", "c", "d"}, l)

	// With middle at the beginning
	la, lb = New(), New()

	la.Append("a")
	la.Append("c")
	lb.Append("b")

	l = la.AppendListMiddle(lb)
	testListOrder(t, []string{"a", "b", "c"}, l)

	// With middle at the end, size 1
	la, lb = New(), New()

	la.Prepend("b")
	la.Prepend("a")
	lb.Append("c")

	l = la.AppendListMiddle(lb)
	testListOrder(t, []string{"a", "b", "c"}, l)
}

func TestListPrependMiddle(t *testing.T) {
	la, lb := New(), New()

	la.Append("e")
	la.Prepend("a")
	lb.Append("b")
	lb.Append("c")
	lb.Append("d")

	l := la.PrependListMiddle(lb)
	testListOrder(t, []string{"a", "b", "c", "d", "e"}, l)

	// Test with list of size 1
	la, lb = New(), New()

	la.Append("c")
	la.Append("d")
	la.Append("e")
	la.Prepend("a")
	lb.Append("b")

	l = la.PrependListMiddle(lb)
	testListOrder(t, []string{"a", "b", "c", "d", "e"}, l)

	// Test empty lists
	la, lb = New(), New()
	la.Append("a")
	la.Append("b")
	la.Append("c")

	l = la.PrependListMiddle(lb)
	testListOrder(t, []string{"a", "b", "c"}, l)

	l = lb.PrependListMiddle(la)
	testListOrder(t, []string{"a", "b", "c"}, l)

	// With middle at the end
	la, lb = New(), New()

	la.Prepend("d")
	la.Prepend("a")
	lb.Append("b")
	lb.Append("c")

	l = la.PrependListMiddle(lb)
	testListOrder(t, []string{"a", "b", "c", "d"}, l)

	// With middle at the beginning
	la, lb = New(), New()

	la.Append("c")
	la.Append("d")
	lb.Append("a")
	lb.Append("b")

	l = la.PrependListMiddle(lb)
	testListOrder(t, []string{"a", "b", "c", "d"}, l)

	// With middle at the beginning, size 1
	la, lb = New(), New()

	la.Append("b")
	la.Append("c")
	lb.Append("a")

	l = la.PrependListMiddle(lb)
	testListOrder(t, []string{"a", "b", "c"}, l)
}


func TestRemove(t *testing.T) {
	l := New()
	a := l.Append("a")
	l.Append("b")
	c := l.Append("c")
	l.Append("d")
	e := l.Append("e")

	// Test deletion of root
	if del, ok := l.Remove(a); !ok || del != "a" {
		t.Errorf("Deleted value incorrect, expected 'a', got '%v'", del)
	}

	testListOrder(t, []string{"b", "c", "d", "e"}, l)

	// Test deletion of end
	if del, ok := l.Remove(e); !ok || del != "e" {
		t.Errorf("Deleted value incorrect, expected 'e', got '%v'", del)
	}

	testListOrder(t, []string{"b", "c", "d"}, l)

	// Test deletion of a middle element
	if del, ok := l.Remove(c); !ok || del != "c" {
		t.Errorf("Deleted value incorrect, expected 'c', got '%v'", del)
	}

	testListOrder(t, []string{"b", "d"}, l)

	// Test removal of element from different list
	la, lb := New(), New()
	la.Append("a")
	if _, ok := la.Remove(lb.Append("b")); ok {
		t.Error("Removal of element from different list did not result in error")
	}
	testListOrder(t, []string{"a"}, la)
	testListOrder(t, []string{"b"}, lb)

	// Test removal of element with nil as value
	lc := New()
	if v, ok := lc.Remove(lc.Append(nil)); !ok || v != nil {
		t.Error("Removal of element from different list did not result in error")
	}

	testListOrder(t, []string{}, lc)
}

func testListOrder(t *testing.T, expected []string, l *List) {
	at := 0
	for e := l.Front(); e != nil; e = e.Next() {
		if at >= len(expected) {
			t.Errorf("List is too long, expected %v, but currently at %v\n", len(expected), at)
		} else if expected[at] != e.Value {
			t.Errorf("Encoundered wrong value in forward iteration, expected %v, but got %v at iteration %d\n", expected[at], e.Value, at)
		}

		at++
	}

	if l.Len() != at || at != len(expected) {
		t.Errorf("List size differs from actual size forwards; found %d, expected %d; list length %d", at, len(expected), l.Len())
	}

	for e := l.End(); e != nil; e = e.Prev() {
		at--

		if expected[at] != e.Value {
			t.Errorf("Encoundered wrong value in backward iteration, expected %v, but got %v at iteration %d\n", expected[at], e.Value, at)
		}
	}

	if 0 != at {
		t.Errorf("List size differs from actual size backwards; found %d, expected %d; list length %d", len(expected)-at, len(expected), l.Len())
	}
}
