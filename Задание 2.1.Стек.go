//стек

package main

import "fmt"

type Stack struct {
    items []int
}

func NewStack() *Stack {
    return &Stack{}
}

func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
    if len(s.items) == 0 {
        return 0, fmt.Errorf("stack is empty")
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, nil
}

func (s *Stack) Peek() (int, error) {
    if len(s.items) == 0 {
        return 0, fmt.Errorf("stack is empty")
    }
    return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}