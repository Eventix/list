package list

import (
	"sync"
)

type (
	Element struct {
		m sync.RWMutex

		Value      interface{}
		prev, next *Element

		l *List
	}

	List struct {
		e   Element
		mid *Element
		m   sync.RWMutex
		len int
	}
)

func (e *Element) Next() *Element {
	e.m.RLock()
	defer e.m.RUnlock()

	return e.next
}

func (e *Element) Prev() *Element {
	e.m.RLock()
	defer e.m.RUnlock()

	return e.prev
}

func (e *Element) Delete() {
	e.m.RLock()
	defer e.m.RUnlock()

	e.l.Remove(e)
}

func New() *List {
	return &List{m : sync.RWMutex{}}
}

func (l *List) Front() *Element {
	l.m.RLock()
	defer l.m.RUnlock()

	return l.e.next
}

func (l *List) End() *Element {
	l.m.RLock()
	defer l.m.RUnlock()

	return l.e.prev
}

func (l *List) Middle() *Element {
	l.m.RLock()
	defer l.m.RUnlock()

	return l.mid
}

func (l *List) Len() int {
	l.m.RLock()
	defer l.m.RUnlock()

	return l.len
}

func (l *List) Append(v interface{}) *Element {
	l.m.RLock()
	p := l.e.prev
	l.m.RUnlock()

	return l.AppendAfter(v, p)
}

func (l *List) AppendAfter(v interface{}, after *Element) *Element {
	ne := &Element{
		Value: v,
		m:     sync.RWMutex{},
		l:     l,
	}

	if l.After(after, ne) {
		l.m.Lock()
		l.len++
		l.m.Unlock()
	}

	return ne
}

func (l *List) After(stays, appended *Element) (s bool) {
	// Ignore cases of wrong list assignments
	if appended == nil || appended.l != l || (stays != nil && stays.l != l) || (l.Len() > 0 && stays == nil) {
		return
	}

	appended.m.Lock()
	defer appended.m.Unlock()

	tail := l.e.prev == stays
	if stays != nil {
		stays.m.Lock()
		defer stays.m.Unlock()

		on := stays.next
		appended.prev = stays
		stays.next = appended
		if on != nil {
			appended.next = on
			on.prev = appended
		}
	} else {
		l.mid = appended
		l.e.next = appended
	}

	if tail {
		l.e.prev = appended
	}

	appended.l = l

	return true
}

func (l *List) Prepend(v interface{}) *Element {
	l.m.RLock()
	p := l.e.next
	l.m.RUnlock()

	return l.PrependBefore(v, p)
}

func (l *List) PrependBefore(v interface{}, before *Element) *Element {
	ne := &Element{
		Value: v,
		m:     sync.RWMutex{},
		l:     l,
	}

	if l.Before(before, ne) {
		l.m.Lock()
		l.len++
		l.m.Unlock()
	}

	return ne
}

func (l *List) Before(stays, before *Element) (s bool) {
	// Ignore cases of wrong list assignments
	if before == nil || before.l != l || (stays != nil && stays.l != l) || (l.Len() > 0 && stays == nil) {
		return
	}

	before.m.Lock()
	defer before.m.Unlock()

	tail := l.e.next == stays
	if stays != nil {
		stays.m.Lock()
		defer stays.m.Unlock()

		on := stays.prev
		before.next = stays
		stays.prev = before
		if on != nil {
			before.prev = on
			on.next = before
		}
	} else {
		l.mid = before
		l.e.prev = before
	}

	if tail {
		l.e.next = before
	}

	before.l = l

	return true
}

func (l *List) AppendList(ll *List) *List {
	l.m.RLock()
	e := l.e.prev
	l.m.RUnlock()

	return l.AppendListAfter(ll, e)
}

func (l *List) AppendListMiddle(ll *List) *List {
	l.m.RLock()
	e := l.mid
	l.m.RUnlock()

	return l.AppendListAfter(ll, e)
}

func (l *List) AppendListAfter(ll *List, e *Element) *List {
	if l == ll {
		return l
	}

	for ne := ll.e.next; ne != nil; ne = ne.next {
		e = l.AppendAfter(ne.Value, e)
	}

	return l
}

func (l *List) PrependList(ll *List) *List {
	l.m.RLock()
	e := l.e.next
	l.m.RUnlock()

	return l.PrependListBefore(ll, e)
}

func (l *List) PrependListMiddle(ll *List) *List {
	l.m.RLock()
	e := l.mid
	l.m.RUnlock()

	return l.PrependListBefore(ll, e)
}
func (l *List) PrependListBefore(ll *List, e *Element) *List {
	if l == ll {
		return l
	}

	for ne := ll.e.prev; ne != nil; ne = ne.prev {
		e = l.PrependBefore(ne.Value, e)
	}

	return l
}

func (l *List) Remove(e *Element) (v interface{}, ok bool) {
	if e.l != l {
		return
	}

	l.m.Lock()
	defer l.m.Unlock()

	v = e.Value
	ok = true

	if l.e.next == e {
		l.e.next = e.next
	}

	if l.e.prev == e {
		l.e.prev = e.prev
	}

	if l.mid == e {
		if e.next == nil {
			l.mid = e.prev
		} else {
			l.mid = e.next
		}
	}

	if e.prev != nil {
		e.prev.next = e.next
	}

	if e.next != nil {
		e.next.prev = e.prev
	}

	l.len--

	return
}
