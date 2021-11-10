package lowest_common_ancestor_of_a_binary_tree

type Node struct {
	Val    int

	Left   *Node
	Right  *Node
	Parent *Node
}

func lowestCommonAncestor(p *Node, q *Node) *Node {
	if p == nil || q == nil {
		return nil
	}

	ap := p
	for ap != nil {
		aq := q
		for aq != nil {
			if ap == aq {
				return aq
			}
			aq = aq.Parent
		}
		ap = ap.Parent
	}
	return nil
}
