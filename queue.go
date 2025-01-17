package collections

import (
	"errors"
	"fmt"
)

type queue[T any] []T

func NewQueue[T any]() queue[T] {
	return []T{}
}

func (s *queue[T]) Len() int {
	return len(*s)
}

func (s *queue[T]) Empty() bool {
	return s.Len() == 0
}

func (s *queue[T]) PushAll(elements queue[T]) {
	*s = append(*s, elements...)
}

func (s *queue[T]) Push(element T) {
	*s = append(*s, element)
}

func (s *queue[T]) Pop() (T, error) {
	element, err := s.Peek()
	if err != nil {
		return element, err
	}
	*s = (*s)[1:]
	return element, nil
}

func (s *queue[T]) PopMultiple(amount int) (stack[T], error) {
	if s.Len() < amount {
		return nil, fmt.Errorf("queue does not contain %v elements", amount)
	}
	result := append(stack[T]{}, (*s)[0:amount]...)
	*s = (*s)[amount:]

	return result, nil
}

func (s *queue[T]) Peek() (T, error) {
	if s.Len() == 0 {
		var empty T
		return empty, errors.New("queue is empty")
	}
	return (*s)[0], nil
}
