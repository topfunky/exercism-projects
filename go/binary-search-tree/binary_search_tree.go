package binarysearchtree

const testVersion = 1

// SearchTreeData stores a single node of the search tree with references to
// left and right nodes. There is no reference back to the parent node.
type SearchTreeData struct {
	data  int
	left  *SearchTreeData
	right *SearchTreeData
}

// Bst creates a pointer to a new SearchTreeData struct.
func Bst(data int) *SearchTreeData {
	node := new(SearchTreeData)
	node.data = data
	return node
}

// Insert adds a node at the correct (ordered) position in the tree.
func (node *SearchTreeData) Insert(data int) {
	switch {
	case data < node.data:
		if node.left == nil {
			node.left = Bst(data)
		} else {
			node.left.Insert(data)
		}
	case data > node.data:
		if node.right == nil {
			node.right = Bst(data)
		} else {
			node.right.Insert(data)
		}
	default:
		// data exists in tree, but make a duplicate (according to test)
		node.left = Bst(data)
	}
}

// MapString runs function f on all nodes in the tree, in ascending order.
// It returns a slice of the resulting transformed values as strings.
func (node *SearchTreeData) MapString(f func(int) string) (results []string) {
	if node.data != 0 {
		if node.left != nil {
			results = append(results, node.left.MapString(f)...)
		}

		results = append(results, f(node.data))

		if node.right != nil {
			results = append(results, node.right.MapString(f)...)
		}
	}

	return results
}

// MapInt runs function f on all nodes in the tree, in ascending order.
// It returns a slice of the resulting transformed values as ints.
func (node *SearchTreeData) MapInt(f func(int) int) (results []int) {
	if node.data != 0 {
		if node.left != nil {
			results = append(results, node.left.MapInt(f)...)
		}

		results = append(results, f(node.data))

		if node.right != nil {
			results = append(results, node.right.MapInt(f)...)
		}
	}

	return results
}
