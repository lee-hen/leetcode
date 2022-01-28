package my_solution

import (
	"container/list"
	"math/rand"
	"time"
)

type Item struct {
	value int
	id int
}

type RandomizedSet struct {
	N          int
	lookup     map[int]*list.Element
	l          *list.List
	elements   []int
	deleted    map[int]*list.Element
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		lookup: make(map[int]*list.Element),
		l: list.New(),
		elements: make([]int, 0),
		deleted: make(map[int]*list.Element),
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.lookup[val]; ok {
		return false
	}

	if r.N >= len(r.lookup) / 2 {
		r.Reset()
	}

	for idx, prev := range r.deleted {
		item := &Item{value: val, id: idx}
		var el *list.Element

		if prev != nil {
			el = r.l.InsertAfter(item, prev)
		} else {
			el = r.l.PushFront(item)
		}

		r.lookup[val] = el
		r.elements[idx] = val
		delete(r.deleted, idx)
		return true
	}

	item := &Item{value: val, id: r.N}
	el := r.l.PushBack(item)
	r.lookup[val] = el
	r.elements = append(r.elements, val)
	r.N++
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	if _, ok := r.lookup[val]; !ok {
		return false
	}

	el := r.lookup[val]
	next := el.Next()
	prev := el.Prev()

	item := r.l.Remove(el).(*Item)

	if idx := item.id; next != nil || prev != nil  {
		if idx < len(r.elements) - 1 && next != nil && r.elements[idx+1] == next.Value.(*Item).value && prev != nil || next == nil {
			r.elements[idx] = prev.Value.(*Item).value
		} else {
			r.elements[idx] = next.Value.(*Item).value
		}
	}

	delete(r.lookup, val)
	r.deleted[item.id] = prev

	if r.N >= len(r.lookup) / 2 {
		r.Reset()
	}
	return true
}

func (r *RandomizedSet) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(r.elements))
	return r.elements[idx]
}

func (r *RandomizedSet) Reset() {
	r.N = 0
	r.elements = make([]int, 0)
	r.deleted = make(map[int]*list.Element)

	curr := r.l.Front()
	for curr != nil {
		item := curr.Value.(*Item)
		item.id = r.N
		r.elements = append(r.elements, item.value)

		curr = curr.Next()
		r.N++
	}
}
