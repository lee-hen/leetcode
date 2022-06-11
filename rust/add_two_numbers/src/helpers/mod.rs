use crate::ListNode;

pub fn add_many(values: Vec<i32>) -> Option<Box<ListNode>> {
    if values.is_empty() {
        return None;
    }

    add_many_recursive(0, values)
}

fn add_many_recursive(i: usize, values: Vec<i32>) -> Option<Box<ListNode>> {
    if i >= values.len() {
        return None;
    }

    let val = values.get(i).unwrap_or(&0);
    let val = *val;

    let more = add_many_recursive(i + 1, values);
    let mut list = Box::new(ListNode::new(val));
    list.next = more;

    Some(list)
}
