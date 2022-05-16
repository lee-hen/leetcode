use std::cell::RefCell;
use std::cmp;
use std::rc::Rc;
use common::*;

pub fn max_path_sum(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    let mut max_path = i32::MIN;
    helper(root, &mut max_path);
    return max_path;
}

fn helper(node: Option<Rc<RefCell<TreeNode>>>, max_path: &mut i32) -> i32 {
    if let Some(node) = node {
        let left = Rc::clone(&node).borrow().left.clone();
        let right = Rc::clone(&node).borrow().right.clone();

        let left_path = helper(left, max_path);
        let right_path = helper(right, max_path);

        let max_with_left_and_right = left_path + right_path + Rc::clone(&node).borrow().val;
        let max_with_left_or_right = cmp::max(
            cmp::max(left_path, right_path) + Rc::clone(&node).borrow().val,
            Rc::clone(&node).borrow().val,
        );
        *max_path = cmp::max(
            cmp::max(*max_path, max_with_left_and_right),
            max_with_left_or_right,
        );
        return max_with_left_or_right;
    }
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_cases() {
        let root = new_tree_node(-10);
        set_left(&root, 9);
        let right = set_right(&root, 20);
        set_left(&right, 15);
        set_right(&right, 7);
        assert_eq!(max_path_sum(root), 42);

        let root = new_tree_node(1);
        set_left(&root, 2);
        assert_eq!(max_path_sum(root), 3);

        let root = new_tree_node(-2);
        set_left(&root, -1);
        set_right(&root, -3);
        assert_eq!(max_path_sum(root), -1);

        let root = new_tree_node(2);
        set_left(&root, -1);
        set_right(&root, -2);
        assert_eq!(max_path_sum(root), 2);

        let root = new_tree_node(9);
        set_left(&root, 6);
        let right = set_right(&root, -3);
        set_left(&right, -6);
        let right = set_right(&right, 2);
        let left = set_left(&right, 2);
        set_right(&left, -6);
        let left = set_left(&left, -6);
        set_left(&left, -6);
        assert_eq!(max_path_sum(root), 16);
    }
}
