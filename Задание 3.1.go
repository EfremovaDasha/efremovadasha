// Стек на дженериках

package main

type Stack[T any] struct {
    items []T
}

func NewStack[T any]() *Stack[T] {
    return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
    if len(s.items) == 0 {
        var zero T
        return zero, fmt.Errorf("stack is empty")
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, nil
}

func (s *Stack[T]) Peek() (T, error) {
    if len(s.items) == 0 {
        var zero T
        return zero, fmt.Errorf("stack is empty")
    }
    return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}