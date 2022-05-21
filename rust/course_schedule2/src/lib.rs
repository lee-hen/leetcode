use std::cell::RefCell;
use std::collections::HashMap;
use std::rc::Rc;

pub fn find_order(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> Vec<i32> {
    let graph = Graph::build_graph(num_courses, prerequisites);

    let mut order = vec![None; num_courses as usize];
    let mut end_of_list = add_non_dependent(&mut order, graph.nodes, 0);
    let mut to_be_processed = 0;

    while to_be_processed < order.len() {
        let current = match order[to_be_processed].as_ref() {
            Some(curr) => Rc::clone(curr),
            None => return Vec::new(),
        };

        let children = Rc::clone(&current).borrow().children.to_owned();
        for child in children.iter() {
            Rc::clone(&child).borrow_mut().dependencies -= 1;
        }
        end_of_list = add_non_dependent(&mut order, children, end_of_list);
        to_be_processed += 1;
    }

    convert(order)
}

fn add_non_dependent(
    order: &mut Vec<Option<Rc<RefCell<Course>>>>,
    courses: Vec<Rc<RefCell<Course>>>,
    offset: usize,
) -> usize {
    let mut offset = offset;
    for course in courses {
        if Rc::clone(&course).borrow().dependencies == 0 {
            order[offset] = Some(Rc::clone(&course));
            offset += 1;
        }
    }
    offset
}

fn convert(courses: Vec<Option<Rc<RefCell<Course>>>>) -> Vec<i32> {
    let mut order = vec![0; courses.len()];
    for i in 0..courses.len() {
        let course = courses[i].clone().unwrap();
        order[i] = Rc::clone(&course).borrow().course_no;
    }
    order
}

struct Course {
    children: Vec<Rc<RefCell<Course>>>,
    visited: HashMap<i32, Rc<RefCell<Course>>>,
    dependencies: i32,
    course_no: i32,
}

impl Course {
    fn new(n: i32) -> Self {
        Self {
            children: Vec::new(),
            visited: HashMap::new(),
            dependencies: 0,
            course_no: n,
        }
    }

    fn add_neighbor(&mut self, node: Rc<RefCell<Course>>) {
        let course_no = Rc::clone(&node).as_ref().borrow().course_no;

        if !self
            .visited
            .contains_key(&course_no)
        {
            self.children.push(Rc::clone(&node));
            self.visited.insert(course_no, Rc::clone(&node));
            Rc::clone(&node).borrow_mut().dependencies += 1;
        }
    }
}

struct Graph {
    nodes: Vec<Rc<RefCell<Course>>>,
    lookup: HashMap<i32, Rc<RefCell<Course>>>,
}

impl Graph {
    fn init() -> Self {
        Self {
            nodes: Vec::new(),
            lookup: HashMap::new(),
        }
    }

    pub fn build_graph(num_courses: i32, dependencies: Vec<Vec<i32>>) -> Self {
        let mut graph = Graph::init();
        for course_no in 0..num_courses {
            let node = Rc::new(RefCell::new(Course::new(course_no)));
            graph.nodes.push(Rc::clone(&node));
            graph.lookup.insert(course_no, Rc::clone(&node));
        }

        for dependency in dependencies {
            let first = dependency[1];
            let second = dependency[0];
            graph.add_edge(first, second);
        }

        graph
    }

    fn add_edge(&self, start: i32, end: i32) {
        let start_node = self.lookup.get(&start).unwrap();
        let end_node = self.lookup.get(&end).unwrap();
        Rc::clone(start_node)
            .borrow_mut()
            .add_neighbor(Rc::clone(end_node));
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_case() {
        assert_eq!(find_order(
            4,
            vec![vec![1, 0], vec![2, 0], vec![3, 1], vec![3, 2]]
        ), vec![0, 1, 2, 3]);
    }
}
