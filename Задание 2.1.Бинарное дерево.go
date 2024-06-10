//бинарное дерево

package main

import "fmt"

type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

type BinaryTree struct {
    Root *TreeNode
}

func NewBinaryTree() *BinaryTree {
    return &BinaryTree{}
}

func (t *BinaryTree) Insert(value int) {
    t.Root = insertRecursive(t.Root, value)
}

func insertRecursive(node *TreeNode, value int) *TreeNode {
    if node == nil {
        return &TreeNode{Value: value}
    }

    if value < node.Value {
        node.Left = insertRecursive(node.Left, value)
    } else {
        node.Right = insertRecursive(node.Right, value)
    }

    return node
}

func (t *BinaryTree) Traverse() {
    traverseInOrder(t.Root)
}

func traverseInOrder(node *TreeNode) {
    if node != nil {
        traverseInOrder(node.Left)
        fmt.Println(node.Value)
        traverseInOrder(node.Right)
    }
}