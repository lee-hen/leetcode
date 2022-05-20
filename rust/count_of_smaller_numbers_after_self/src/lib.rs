use std::marker::Copy;

#[derive(Copy, Clone)]
struct Element {
    val: i32,
    idx: usize,
}

impl Element {
    pub fn new(val: i32, idx: usize) -> Self {
        Self { val, idx }
    }
}

pub fn count_smaller(nums: Vec<i32>) -> Vec<i32> {
    let mut counts: Vec<i32> = vec![0; nums.len()];
    let mut main_array: Vec<Element> = Vec::new();
    for i in 0..nums.len() {
        main_array.push(Element::new(nums[i], i))
    }
    let mut auxiliary_array: Vec<Element> = Vec::new();
    copy(&mut auxiliary_array, &main_array);

    merge_sort_helper(
        &mut main_array,
        &mut auxiliary_array,
        &mut counts,
        0,
        nums.len() - 1,
    );
    counts
}

fn merge_sort_helper(
    main_array: &mut Vec<Element>,
    auxiliary_array: &mut Vec<Element>,
    counts: &mut Vec<i32>,
    start_idx: usize,
    end_idx: usize,
) {
    if start_idx == end_idx {
        return;
    }

    let middle_idx = start_idx + (end_idx - start_idx) / 2;
    merge_sort_helper(auxiliary_array, main_array, counts, start_idx, middle_idx);
    merge_sort_helper(auxiliary_array, main_array, counts, middle_idx + 1, end_idx);

    do_merge(
        main_array,
        auxiliary_array,
        counts,
        start_idx,
        middle_idx,
        end_idx,
    );
}

fn do_merge(
    main_array: &mut Vec<Element>,
    auxiliary_array: &mut Vec<Element>,
    counts: &mut Vec<i32>,
    start_idx: usize,
    middle_idx: usize,
    end_idx: usize,
) {
    let mut k = start_idx;
    let mut i = start_idx;
    let mut j = middle_idx + 1;

    while i <= middle_idx && j <= end_idx {
        if auxiliary_array[i].val <= auxiliary_array[j].val {
            // When we select an element i on the left array,
            // we know that elements selected previously from the right array jump from i's right to i's left.
            // For each element i, records the number of elements jumping from i's right to i's left during the merge sort.
            counts[auxiliary_array[i].idx] += (j - middle_idx - 1) as i32;
            main_array[k] = auxiliary_array[i];
            i += 1;
        } else {
            main_array[k] = auxiliary_array[j];
            j += 1; // when j increased, j-middleIdx-1 is the number of elements jumping from i's right to i's left.
        }
        k += 1;
    }

    while i <= middle_idx {
        counts[auxiliary_array[i].idx] += (j - middle_idx - 1) as i32;
        main_array[k] = auxiliary_array[i];
        i += 1;
        k += 1;
    }

    while j <= end_idx {
        main_array[k] = auxiliary_array[j];
        j += 1;
        k += 1;
    }
}

fn copy(auxiliary_array: &mut Vec<Element>, main_array: &Vec<Element>) {
    for el in main_array {
        auxiliary_array.push(Element::new(el.val, el.idx))
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_cases() {
        let nums = vec![5, 2, 6, 1];
        assert_eq!(count_smaller(nums), vec![2, 1, 1, 0]);

        let nums = vec![8, 5, 2, 9, 5, 6, 3];
        assert_eq!(count_smaller(nums), vec![5, 2, 0, 3, 1, 1, 0]);

        let nums = vec![11, 9, 6, 7, 4, 5, 8, 20, 12];
        assert_eq!(count_smaller(nums), vec![6, 5, 2, 2, 0, 0, 0, 1, 0]);

        let nums = vec![5, 2, 6, 1, 4, 9, 3, 8, 7, 6, 4, 5, 1];
        assert_eq!(
            count_smaller(nums),
            vec![6, 2, 6, 0, 2, 7, 1, 5, 4, 3, 1, 1, 0]
        );
    }
}
