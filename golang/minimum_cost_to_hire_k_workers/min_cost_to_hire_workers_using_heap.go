package minimum_cost_to_hire_k_workers

import (
	"container/heap"
	"math"
	"sort"
)

func MinCostToHireWorkersHeapImpl(quality []int, wage []int, k int) float64 {
	candidates := make([]*candidate, 0)
	for i := range wage {
		candidates = append(candidates, &candidate{quality[i], float64(wage[i]), float64(wage[i]) / float64(quality[i])})
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].ratio < candidates[j].ratio
	})

	var minCost = math.MaxFloat64
	sumQuality := 0
	pool := make(MaxHeap, 0)
	heap.Init(&pool)

	for _, candidate := range candidates {
		heap.Push(&pool, candidate.quality)
		sumQuality += candidate.quality

		if pool.Len() > k {
			sumQuality -= heap.Pop(&pool).(int)
		}

		if pool.Len() == k {
			minCost = math.Min(minCost, float64(sumQuality) * candidate.ratio)
		}
	}

	return minCost
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
