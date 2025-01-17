package collections

import (
	"errors"
	"fmt"
)

type stack[T any] []T

func NewStack[T any]() stack[T] {
	return []T{}
}

func (s *stack[T]) Len() int {
	return len(*s)
}

func (s *stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *stack[T]) PushAll(elements stack[T]) {
	*s = append(elements, *s...)
}

func (s *stack[T]) Push(element T) {
	*s = append([]T{element}, *s...)
}

func (s *stack[T]) Pop() (T, error) {
	element, err := s.Peek()
	if err != nil {
		return element, err
	}
	*s = (*s)[1:]
	return element, nil
}

func (s *stack[T]) PopMultiple(amount int) (stack[T], error) {
	if s.Len() < amount {
		return nil, fmt.Errorf("stack does not contain %v elements", amount)
	}
	result := append(stack[T]{}, (*s)[0:amount]...)
	*s = (*s)[amount:]

	return result, nil
}

func (s *stack[T]) Peek() (T, error) {
	if s.Len() == 0 {
		var empty T
		return empty, errors.New("stack is empty")
	}
	return (*s)[0], nil
}
