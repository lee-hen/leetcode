// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}

impl Default for ListNode {
  fn default() -> ListNode {
    ListNode{
      next: None,
      val: 0,
    }
  }
}

pub fn add_two_numbers(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
  fn add_lists(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>, carry: i32) -> Option<Box<ListNode>> {
    if carry == 0 && l1 == None && l2 == None {
      return None;
    }

    let mut value = carry;

    let l1 = l1.unwrap_or_default();
    let l2 = l2.unwrap_or_default();

    value += l1.val;
    value += l2.val;

    let next1 = l1.next;
    let next2 = l2.next;

    let more = match value {
      v if v >= 10 => add_lists(next1, next2, 1),
      _ => add_lists(next1, next2, 0),
    };

    let mut list_node = ListNode::new(value % 10);
    list_node.next = more;

    Some(Box::new(list_node))
  }

  add_lists(l1, l2, 0)
}

pub mod helpers;

#[cfg(test)]
mod tests {
  use crate::helpers::add_many;
  use super::*;

  #[test]
  fn test_case() {
    let l1 = vec![9, 9, 9, 9, 9, 9, 9];
    let l2 = vec![9, 9, 9, 9];
    let l1 = add_many(l1);
    let l2 = add_many(l2);
    let expected = vec![8, 9, 9, 9, 0, 0, 0, 1];
    let expected = add_many(expected);
    assert_eq!(expected, add_two_numbers(l1, l2));
  }
}
