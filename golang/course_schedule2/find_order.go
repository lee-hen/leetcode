package course_schedule2

func FindOrder(numCourses int, prerequisites [][]int) []int {
	graph := buildGraph(numCourses, prerequisites)

	order := make([]*Course, numCourses)
	endOfList := addNonDependent(order, graph.nodes, 0)
	toBeProcessed := 0
	for toBeProcessed < len(order) {
		current := order[toBeProcessed]
		if current == nil {
			return []int{}
		}

		children := current.children
		for _, child := range children {
			child.dependencies--
		}

		endOfList = addNonDependent(order, children, endOfList)
		toBeProcessed++
	}

	return convert(order)
}

func convert(courses []*Course) []int {
	order := make([]int, len(courses))
	for i := 0; i < len(courses); i++ {
		order[i] = courses[i].courseNo
	}
	return order
}

func addNonDependent(order, courses []*Course, offset int) int {
	for _, course := range courses {
		if course.dependencies == 0 {
			order[offset] = course
			offset++
		}
	}
	return offset
}

type Course struct {
	children     []*Course
	visited      map[int]*Course
	dependencies int
	courseNo     int
}

func NewCourse(n int) *Course {
	p := new(Course)
	p.courseNo = n
	p.visited = make(map[int]*Course)
	return p
}

func (p *Course) AddNeighbor(node *Course) {
	if _, ok := p.visited[node.courseNo]; !ok {
		p.children = append(p.children, node)
		p.visited[node.courseNo] = node
		node.dependencies++
	}
}

type Graph struct {
	nodes []*Course
	lookup map[int]*Course
}

func buildGraph(numCourses int, dependencies [][]int) *Graph {
	graph := new(Graph)
	graph.lookup = make(map[int]*Course)

	for courseNo := 0; courseNo < numCourses; courseNo++ {
		node := NewCourse(courseNo)
		graph.nodes = append(graph.nodes, node)
		graph.lookup[courseNo] = node
	}

	for _, dependency := range dependencies {
		first, second := dependency[1], dependency[0]
		graph.AddEdge(first, second)
	}
	return graph
}

func (g *Graph) AddEdge(start, end int) {
	startNode := g.lookup[start]
	endNode := g.lookup[end]

	startNode.AddNeighbor(endNode)
}
