// Бинарное дерево на дженериках

package main

type TreeNode[T any] struct {
    Value T
    Left  *TreeNode[T]
    Right *TreeNode[T]
}

type BinaryTree[T any] struct {
    Root *TreeNode[T]
}

func NewBinaryTree[T any]() *BinaryTree[T] {
    return &BinaryTree[T]{}
}

func (t *BinaryTree[T]) Insert(value T) {
    t.Root = insertRecursive(t.Root, value)
}

func insertRecursive[T any](node *TreeNode[T], value T) *TreeNode[T] {
    if node == nil {
        return &TreeNode[T]{Value: value}
    }

    if value.(int) < node.Value.(int) {
        node.Left = insertRecursive(node.Left, value)
    } else {
        node.Right = insertRecursive(node.Right, value)
    }

    return node
}

func (t *BinaryTree[T]) Traverse() {
    traverseInOrder(t.Root)
}

func traverseInOrder[T any](node *TreeNode[T]) {
    if node != nil {
        traverseInOrder(node.Left)
        fmt.Println(node.Value)
        traverseInOrder(node.Right)
    }
}