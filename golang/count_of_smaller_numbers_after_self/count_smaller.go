package count_of_smaller_numbers_after_self

type element struct {
	val int
	idx int
}

func CountSmaller(nums []int) []int {
	counts := make([]int, len(nums))
	mainArray := make([]*element, 0)
	for i := range nums {
		mainArray = append(mainArray, &element{nums[i], i})
	}

	auxiliaryArray := make([]*element, len(nums))
	copy(auxiliaryArray, mainArray)
	mergeSortHelper(mainArray, auxiliaryArray, counts, 0, len(nums)-1)
	return counts
}

func mergeSortHelper(mainArray, auxiliaryArray []*element, counts []int, startIdx, endIdx int) {
	if startIdx == endIdx {
		return
	}
	middleIdx := (startIdx + endIdx) / 2
	// sort left auxiliary array
	mergeSortHelper(auxiliaryArray, mainArray, counts, startIdx, middleIdx)
	// sort right auxiliary array
	mergeSortHelper(auxiliaryArray, mainArray, counts, middleIdx+1, endIdx)

	// My solution, Accepted.
	//hi := endIdx
	//mid := middleIdx
	//
	//for mid >= startIdx && hi > middleIdx {
	//	if auxiliaryArray[mid].val > auxiliaryArray[hi].val {
	//		counts[auxiliaryArray[mid].idx] += hi-middleIdx
	//		mid--
	//	} else {
	//		hi--
	//	}
	//}

	// then do merge them back to the main array
	doMerge(mainArray, auxiliaryArray, counts, startIdx, middleIdx, endIdx)
}

func doMerge(mainArray, auxiliaryArray []*element, counts []int, startIdx, middleIdx, endIdx int) {
	k := startIdx
	i := startIdx
	j := middleIdx + 1

	for i <= middleIdx && j <= endIdx {
		if auxiliaryArray[i].val <= auxiliaryArray[j].val {
			// When we select an element i on the left array,
			// we know that elements selected previously from the right array jump from i's right to i's left.
			// For each element i, records the number of elements jumping from i's right to i's left during the merge sort.
			counts[auxiliaryArray[i].idx] += j-middleIdx-1
			mainArray[k] = auxiliaryArray[i]
			i++
		} else {
			mainArray[k] = auxiliaryArray[j]
			j++ // when j increased, j-middleIdx-1 is the number of elements jumping from i's right to i's left.
		}
		k++
	}

	for i <= middleIdx {
		counts[auxiliaryArray[i].idx] += j-middleIdx-1
		mainArray[k] = auxiliaryArray[i]
		i++
		k++
	}

	for j <= endIdx {
		mainArray[k] = auxiliaryArray[j]
		j++
		k++
	}
}

func countSmaller(nums []int) []int {
	counts := make([]int, len(nums))

	for i := range nums {
		for j := i+1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				counts[i]++
			}
		}
	}

	return counts
}
