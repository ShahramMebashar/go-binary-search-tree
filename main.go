package main

type BinaryTree[T int | float32] struct {
	Value T
	Left  *BinaryTree[T]
	Right *BinaryTree[T]
}

func insert[T int | float32](node *BinaryTree[T], val T) *BinaryTree[T] {
	if node == nil {
		return &BinaryTree[T]{
			Value: val,
		}
	}
	if val > node.Value {
		node.Right = insert[T](node.Right, val)
	}

	if val < node.Value {
		node.Left = insert[T](node.Left, val)
	}
	return node
}

func makeBtree[T int | float32](mp []T) *BinaryTree[T] {
	var node *BinaryTree[T]
	for _, v := range mp {
		node = insert[T](node, v)
	}

	return node
}

func inOrder[T int | float32](btree *BinaryTree[T]) {
	if btree == nil {
		return
	}
	inOrder[T](btree.Left)
	println(btree.Value)
	inOrder[T](btree.Right)
}

func postOrder[T int | float32](btree *BinaryTree[T]) {
	if btree == nil {
		return
	}
	println(btree.Value)
	postOrder[T](btree.Left)
	postOrder[T](btree.Right)
}

func preOrder[T int | float32](btree *BinaryTree[T]) {
	if btree == nil {
		return
	}
	preOrder[T](btree.Left)
	preOrder[T](btree.Right)
	println(btree.Value)
}

func main() {
	// btree := &BinaryTree[int]{Value: 1}
	btree := makeBtree[int]([]int{3, 6, 7, 10, 2, 1, 4})

	println("Inorder:")
	inOrder[int](btree)

	println("Post order:")
	postOrder[int](btree)

	println("Pre order:")
	preOrder[int](btree)
}
