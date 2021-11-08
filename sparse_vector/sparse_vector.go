package sparse_vector

type SparseVector struct {
	nums map[int]int
}

func Constructor(nums []int) SparseVector {
	v := new(SparseVector)
	v.nums = make(map[int]int)
	for i, num := range nums {
		if num != 0 {
			v.nums[i] = num
		}
	}

	return *v
}

// Return the dotProduct of two sparse vectors
func (v *SparseVector) dotProduct(o SparseVector) int {
	if len(v.nums) == 0 || len(o.nums) == 0 {
		return 0
	}

	var product int
	for i := range v.nums {
		if _, ok := o.nums[i]; ok {
			product += v.nums[i] * o.nums[i]
		}
	}

	return product
}
