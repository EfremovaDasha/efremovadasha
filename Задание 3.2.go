// Двусвязанный список на дженериках

package main

type Node[T any] struct {
    Value T
    Prev  *Node[T]
    Next  *Node[T]
}

type LinkedList[T any] struct {
    Head *Node[T]
    Tail *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
    return &LinkedList[T]{}
}

func (l *LinkedList[T]) AddFront(value T) {
    newNode := &Node[T]{Value: value, Next: l.Head}
    if l.Head != nil {
        l.Head.Prev = newNode
    }
    l.Head = newNode
    if l.Tail == nil {
        l.Tail = newNode
    }
}

func (l *LinkedList[T]) AddBack(value T) {
    newNode := &Node[T]{Value: value, Prev: l.Tail}
    if l.Tail != nil {
        l.Tail.Next = newNode
    }
    l.Tail = newNode
    if l.Head == nil {
        l.Head = newNode
    }
}

func (l *LinkedList[T]) RemoveFront() (T, error) {
    if l.Head == nil {
        var zero T
        return zero, fmt.Errorf("list is empty")
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

func (l *LinkedList[T]) RemoveBack() (T, error) {
    if l.Tail == nil {
        var zero T
        return zero, fmt.Errorf("list is empty")
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