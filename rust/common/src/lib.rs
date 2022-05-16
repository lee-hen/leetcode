use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}


pub fn set_left(node: &Option<Rc<RefCell<TreeNode>>>, val: i32) -> Option<Rc<RefCell<TreeNode>>> {
    let node = Rc::clone(node.as_ref().unwrap());
    let left_node = Rc::clone(&new_tree_node(val).unwrap());
    node.borrow_mut().left = Some(Rc::clone(&left_node));
    left_node.into()
}

pub fn set_right(node: &Option<Rc<RefCell<TreeNode>>>, val: i32) -> Option<Rc<RefCell<TreeNode>>> {
    let node = Rc::clone(node.as_ref().unwrap());
    let left_node = Rc::clone(&new_tree_node(val).unwrap());
    node.borrow_mut().right = Some(Rc::clone(&left_node));
    left_node.into()
}

pub fn new_tree_node(val: i32) -> Option<Rc<RefCell<TreeNode>>> {
    Some(Rc::new(RefCell::new(TreeNode {
        val,
        left: None,
        right: None,
    })))
}
