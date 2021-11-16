package nested_list_weight_sum2

import (
	"container/list"
	"math"
)

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

func DepthSumInverse(nestedList []*NestedInteger) int {
	q := list.New()

	for _, l := range nestedList {
		q.PushBack(l)
	}

	depth := 1
	maxDepth := 0
	sumOfElements := 0
	sumOfProducts := 0

	for q.Len() > 0 {
		size := q.Len()
		maxDepth = max(maxDepth, depth)

		for i := 0; i < size; i++ {
			front := q.Front()
			q.Remove(front)
			nested := front.Value.(*NestedInteger)


			if nested.IsInteger() {
				sumOfElements += nested.GetInteger()
				sumOfProducts += nested.GetInteger() * depth
			} else {
				for _, l := range nested.GetList() {
					q.PushBack(l)
				}
			}
		}
		depth++
	}
	return (maxDepth + 1) * sumOfElements - sumOfProducts
}

func depthSumInverse(nestedList []*NestedInteger) int {
	depths := make(map[*NestedInteger]int)
	maxDepth := getDepth(nestedList, 0, depths)

	var sum int
	for l, depth := range depths {
		sum += (maxDepth - depth + 1) * l.GetInteger()
	}
	return sum
}

func getDepth(nestedList []*NestedInteger, depth int, depths map[*NestedInteger]int) int {
	maxDepth := math.MinInt32
	for _, l := range nestedList {
		if l.IsInteger() {
			depths[l] = depth+1
			maxDepth = max(maxDepth, depths[l])
		} else {
			maxDepth = max(maxDepth, getDepth(l.GetList(), depth+1, depths))
		}
	}
	return maxDepth
}

func max(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr < num {
			curr = num
		}
	}
	return curr
}
