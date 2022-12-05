package gocollections

import (
	"errors"
	"strconv"
)

type stack[T any] []T

func CreateNew[T any]() stack[T] {
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
		return nil, errors.New("stack does not contain " + strconv.Itoa(amount) + " elements")
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
