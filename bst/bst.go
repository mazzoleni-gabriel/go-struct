package bst

type node struct {
	data  int
	right *node
	left  *node
}

type BST struct {
	root *node
}

func (bst *BST) Insert(data int) {
	bst.root = insertNode(bst.root, data)
}

func insertNode(root *node, data int) *node {
	if root == nil {
		return &node{data: data}
	}

	if data <= root.data {
		root.left = insertNode(root.left, data)
		return root
	}

	if data > root.data {
		root.right = insertNode(root.right, data)
		return root
	}
	return nil
}

func (bst *BST) Search(data int) bool {
	return search(bst.root, data) != nil
}

func search(root *node, data int) *node {
	if root == nil {
		return nil
	}

	if data < root.data {
		return search(root.left, data)
	}

	if data > root.data {
		return search(root.right, data)
	}

	return root
}
