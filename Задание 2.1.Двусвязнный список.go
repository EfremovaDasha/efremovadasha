//двусвязнный список

package main

import "fmt"

type Node struct {
    Value int
    Prev  *Node
    Next  *Node
}

type LinkedList struct {
    Head *Node
    Tail *Node
}

func NewLinkedList() *LinkedList {
    return &LinkedList{}
}

func (l *LinkedList) AddFront(value int) {
    newNode := &Node{Value: value, Next: l.Head}
    if l.Head != nil {
        l.Head.Prev = newNode
    }
    l.Head = newNode
    if l.Tail == nil {
        l.Tail = newNode
    }
}

func (l *LinkedList) AddBack(value int) {
    newNode := &Node{Value: value, Prev: l.Tail}
    if l.Tail != nil {
        l.Tail.Next = newNode
    }
    l.Tail = newNode
    if l.Head == nil {
        l.Head = newNode
    }
}

func (l *LinkedList) RemoveFront() (int, error) {
    if l.Head == nil {
        return 0, fmt.Errorf("list is empty")
    }
    value := l.Head.Value
    l.Head = l.Head.Next
    if l.Head != nil {
        l.Head.Prev = nil
    } else {
        l.Tail = nil
    }
    return value, nil
}

func (l *LinkedList) RemoveBack() (int, error) {
    if l.Tail == nil {
        return 0, fmt.Errorf("list is empty")
    }
    value := l.Tail.Value
    l.Tail = l.Tail.Prev
    if l.Tail != nil {
        l.Tail.Next = nil
    } else {
        l.Head = nil
    }
    return value, nil
}