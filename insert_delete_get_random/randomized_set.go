package insert_delete_get_random

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	lookup map[int]int
	elements []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		lookup: make(map[int]int),
		elements: make([]int, 0),
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.lookup[val]; ok {
		return false
	}

	r.lookup[val] = len(r.elements)
	r.elements = append(r.elements, val)
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	if _, ok := r.lookup[val]; !ok {
		return false
	}

	// move the last element to the place idx of the element to delete
	lastElement := r.elements[len(r.elements)-1]
	idx := r.lookup[val]
	r.elements[idx] = lastElement
	r.lookup[lastElement] = idx

	// delete the last element
	r.elements = r.elements[:len(r.elements)-1]
	delete(r.lookup, val)
	return true
}

func (r *RandomizedSet) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(r.elements))
	return r.elements[idx]
}
