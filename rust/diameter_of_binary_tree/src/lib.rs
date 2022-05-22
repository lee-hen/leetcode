use common::*;
use std::rc::Rc;
use std::cell::RefCell;

pub fn diameter_of_binary_tree(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    let mut longest_path = i32::MIN;
    helper(root, &mut longest_path);

    longest_path
}

fn helper(t: Option<Rc<RefCell<TreeNode>>>, longest_path: &mut i32) -> i32 {
    match t {
        Some(t) => {
            let left = t.borrow().left.clone();
            let right = t.borrow().right.clone();

            let left_height = helper(left, longest_path);
            let right_height = helper(right, longest_path);
            *longest_path = i32::max(longest_path.clone(), left_height + right_height);
            return i32::max(left_height, right_height) + 1
        },
        None => 0,
    }
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
        set_right(&root, 3);
        assert_eq!(3, diameter_of_binary_tree(root));

        let root = new_tree_node(1);
        set_left(&root, 2);
        assert_eq!(1, diameter_of_binary_tree(root));
    }
}
