use common::*;
use std::cell::RefCell;
use std::cmp;
use std::collections::{HashMap, VecDeque};
use std::rc::Rc;

pub fn vertical_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
    if let Some(root) = root {
        let mut vertical_orders: HashMap<i32, Vec<i32>> = HashMap::new();
        bfs(&root, &mut vertical_orders);
        return convert_vertical_orders(&vertical_orders);
    }
    Vec::new()
}

fn bfs(root: &Rc<RefCell<TreeNode>>, vertical_orders: &mut HashMap<i32, Vec<i32>>) {
    let mut queue = VecDeque::new();
    let mut curr_indices = VecDeque::new();

    queue.push_back(root.clone());
    curr_indices.push_back(0);

    while !queue.is_empty() {
        let parent = queue.pop_front().unwrap();
        let parent_idx = curr_indices.pop_front().unwrap();

        let vertical_order = vertical_orders
            .entry(parent_idx)
            .or_insert_with(|| vec![Rc::clone(&parent).borrow().val]);
        vertical_order.push(Rc::clone(&parent).borrow().val);

        if Rc::clone(&parent).borrow().left.is_some() {
            queue.push_back(Rc::clone(&parent).borrow().left.as_ref().unwrap().clone());
            curr_indices.push_back(parent_idx - 1);
        }

        if Rc::clone(&parent).borrow().right.is_some() {
            queue.push_back(Rc::clone(&parent).borrow().right.as_ref().unwrap().clone());
            curr_indices.push_back(parent_idx + 1);
        }
    }
}

fn convert_vertical_orders(vertical_orders: &HashMap<i32, Vec<i32>>) -> Vec<Vec<i32>> {
    let mut min_idx = i32::MAX;
    let mut max_idx = i32::MIN;

    for key in vertical_orders.clone().into_keys() {
        min_idx = cmp::min(min_idx, key);
        max_idx = cmp::max(max_idx, key);
    }

    let mut result = Vec::new();
    let mut key = min_idx;
    while key <= max_idx {
        result.push(vertical_orders.get(&key).unwrap().clone());

        key += 1;
    }
    result
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_case1() {
        let root = new_tree_node(3);
        set_left(&root, 9);
        let right = set_right(&root, 20);
        set_left(&right, 15);
        set_right(&right, 7);
        assert_eq!(
            vec![vec![9], vec![3, 15], vec![20], vec![7]],
            vertical_order(root)
        );
    }

    #[test]
    fn test_case2() {
        let root = new_tree_node(3);
        let left = set_left(&root, 9);
        set_left(&left, 4);
        set_right(&left, 0);
        let right = set_right(&root, 8);
        set_left(&right, 1);
        set_right(&right, 7);
        assert_eq!(
            vec![vec![4], vec![9], vec![3, 0, 1], vec![8], vec![7]],
            vertical_order(root)
        );
    }

    #[test]
    fn test_case3() {
        let root = new_tree_node(3);
        let left = set_left(&root, 9);
        set_left(&left, 4);
        let right = set_right(&left, 0);
        set_left(&right, 5);
        set_right(&right, 2);

        let right = set_right(&root, 8);
        set_left(&right, 1);
        set_right(&right, 7);
        assert_eq!(
            vec![vec![4], vec![9, 5], vec![3, 0, 1], vec![8, 2], vec![7]],
            vertical_order(root)
        );
    }

    #[test]
    fn test_case4() {
        let vec: Vec<Vec<i32>> = Vec::new();
        assert_eq!(vec, vertical_order(None));
    }
}
