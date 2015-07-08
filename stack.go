package main

import (
	"fmt"
)

type Stack struct {
	top *Element
	size int
}

type Element struct {
	value interface{}
	next *Element
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}


func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value = s.top.value
		s.top = s.top.next
		s.size--
		return
	}
	return nil
}

func main() {
	stack := new (Stack)
	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	fmt.Printf("stack size %d\n", stack.size)
}
