package next_permutation

func NextPermutation(nums []int)  {
	//Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
	var k = -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			k = i
		}
	}

	if k >= 0 {
		// Find the largest index l greater than k such that a[l] > a[k].
		var l = -1
		for i := 0; i < len(nums); i++ {
			if nums[k] < nums[i] {
				l = i // index l will never smaller than index k
			}
		}

		//Swap the value of a[k] with that of a[l].
		if l != -1 {
			nums[k], nums[l] = nums[l], nums[k]
		}
	}

	//Reverse the sequence from a[k + 1] up to and including the final element a[n].
	ReverseSlice(nums, k+1, len(nums)-1)
}

func ReverseSlice(slice []int, start, end int) {
	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start += 1
		end -= 1
	}
}
