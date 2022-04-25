package nested_list_weight_sum

import "container/list"

// NestedInteger
// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	value int
	nestedList []*NestedInteger
	isInteger bool
}

// IsInteger
// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return n.isInteger
}

// GetInteger
// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return n.value
}

// GetList
// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return n.nestedList
}

// SetInteger
// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.value = value
}

func (n *NestedInteger) Add(elem NestedInteger) {
	n.nestedList = append(n.nestedList, &elem)
}

func depthSum(nestedList []*NestedInteger) int {
	q := list.New()

	for _, l := range nestedList {
		q.PushBack(l)
	}

	depth := 1
	sumOfProducts := 0

	for q.Len() > 0 {
		size := q.Len()

		for i := 0; i < size; i++ {
			front := q.Front()
			q.Remove(front)
			nested := front.Value.(*NestedInteger)


			if nested.IsInteger() {
				sumOfProducts += nested.GetInteger() * depth
			} else {
				for _, l := range nested.GetList() {
					q.PushBack(l)
				}
			}
		}
		depth++
	}
	return sumOfProducts
}
