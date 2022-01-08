package minimum_cost_to_hire_k_workers

import (
	"fmt"
	"math"
	"sort"
)

type candidate struct {
	quality int
	wage float64
	ratio float64
}

func (c *candidate) String() string {
	return fmt.Sprintf("{quality: %d, wage:%f, ratio:%f}\n", c.quality, c.wage, c.ratio)
}

func MinCostToHireWorkers(quality []int, wage []int, k int) float64 {
	candidates := make([]*candidate, 0)
	for i := range wage {
		candidates = append(candidates, &candidate{quality[i], float64(wage[i]), float64(wage[i]) / float64(quality[i])})
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].ratio > candidates[j].ratio
	})

	cpyCandidates := make([]*candidate, len(candidates))
	copy(cpyCandidates, candidates)
	sort.Slice(cpyCandidates, func(i, j int) bool {
		return cpyCandidates[i].quality < cpyCandidates[j].quality
	})

	indices := make(map[*candidate]int)
	for i := range cpyCandidates {
		indices[cpyCandidates[i]] = i
 	}

 	mark := make(map[*candidate]bool)
	var minCost = math.MaxFloat64
	for i := 0; i < len(candidates)-k+1; i++ {
		idx := indices[candidates[i]]
		var cnt = 1

		var currCost = candidates[i].wage
		var ratio = candidates[i].ratio

		for left := idx-1; left >= 0 && cnt < k; left-- {
			if mark[cpyCandidates[left]] {
				continue
			}
			currCost += ratio * float64(cpyCandidates[left].quality)
			cnt++
		}

		for right := idx+1; right < len(cpyCandidates) && cnt < k; right++ {
			if mark[cpyCandidates[right]] {
				continue
			}
			currCost += ratio * float64(cpyCandidates[right].quality)
			cnt++
		}

		if cnt == k {
			minCost = math.Min(currCost, minCost)
		}
		mark[candidates[i]] = true
	}

	return minCost
}
