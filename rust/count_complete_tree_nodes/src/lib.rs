use common::*;
use std::cell::RefCell;
use std::rc::Rc;

pub fn count_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    if root.is_none() {
        return 0;
    }

    let d = depth_without_last_level(&root);
    if d == 0 {
        return 1;
    }

    let mut left = 1;
    let mut right = i32::pow(2, d as u32) - 1;

    while left <= right {
        let pivot = left + (right - left) / 2;
        if exists(pivot, d, &root) {
            left = pivot + 1
        } else {
            right = pivot - 1
        }
    }

    let last_level_nodes = left;
    i32::pow(2, d as u32) - 1 + last_level_nodes
}

fn depth_without_last_level(node: &Option<Rc<RefCell<TreeNode>>>) -> i32 {
    let mut d = 0;
    let mut node = node.clone();
    while let Some(left) = node.clone().unwrap().borrow().left.clone() {
        node = Some(left);
        d += 1;
    }
    d
}

fn exists(target_idx: i32, d: i32, node: &Option<Rc<RefCell<TreeNode>>>) -> bool {
    let mut left = 0;
    let mut right = i32::pow(2, d as u32) - 1;
    let mut node = node.clone();

    for _ in 0..d {
        let pivot = left + (right - left) / 2;

        if target_idx <= pivot {
            node = match node {
                Some(n) => n.borrow().left.clone(),
                None => return false,
            };

            right = pivot;
        } else {
            node = match node {
                Some(n) => n.borrow().right.clone(),
                None => return false,
            };
            left = pivot + 1;
        }
    }
    node.is_some()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_case() {
        let root = new_tree_node(1);
        let left = set_left(&root, 2);
        set_left(&left, 4);
        set_right(&left, 5);

        let right = set_right(&root, 3);
        set_left(&right, 6);
        assert_eq!(count_nodes(root), 6);
    }
}
