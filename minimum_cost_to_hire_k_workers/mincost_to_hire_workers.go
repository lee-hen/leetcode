package minimum_cost_to_hire_k_workers

import (
	"container/heap"
	"math"
)

func minCostToHireWorkers(quality []int, wage []int, k int) float64 {
	var minCost = math.MaxFloat64
	for i := range wage {
		ratio := float64(wage[i]) / float64(quality[i])

		costs := &CostHeap{float64(wage[i])}
		heap.Init(costs)

		for j := range wage {
			if i != j && ratio * float64(quality[j]) >= float64(wage[j]) {
				heap.Push(costs, ratio * float64(quality[j]))
			}
		}

		if costs.Len() < k {
			continue
		}

		var cost float64
		var cnt int
		for costs.Len() > 0 && cnt < k  {
			cost += heap.Pop(costs).(float64)
			cnt++
		}
		minCost = math.Min(cost, minCost)
	}

	return minCost
}

type CostHeap []float64

func (h CostHeap) Len() int           { return len(h) }
func (h CostHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h CostHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *CostHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *CostHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
