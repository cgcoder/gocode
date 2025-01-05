package util

import (
	"errors"
)

type Stack[T any] interface {
	Push(data T)
	Pop() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	Size() int
	UpdateTop(data T) error
}

type stack[T any] struct {
	data []T
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{
		make([]T, 0),
	}
}

func (s *stack[T]) Push(data T) {
	s.data = append(s.data, data)
}

func (s *stack[T]) Pop() (T, error) {
	if len(s.data) == 0 {
		return *new(T), errors.New("stack is empty")
	}
	data := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return data, nil
}

func (s *stack[T]) Peek() (T, error) {
	if len(s.data) == 0 {
		return *new(T), errors.New("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func (s *stack[T]) UpdateTop(data T) error {
	if len(s.data) == 0 {
		return errors.New("stack is empty")
	}
	s.data[len(s.data)-1] = data
	return nil
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *stack[T]) Size() int {
	return len(s.data)
}
