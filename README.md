# Binary Search Tree in Go

Simple Binary Search tree implementation in Golan with generics

```go
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
```
