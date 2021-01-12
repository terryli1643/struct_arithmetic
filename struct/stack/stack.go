package stack

import (
	"errors"
)

type MyStack struct {
	data []interface{}
	top  int
}

func NewMyStack(cap int) *MyStack {
	if cap <= 0 {
		return nil
	}

	return &MyStack{
		data: make([]interface{}, cap),
		top:  -1,
	}
}

func (s *MyStack) Push(e interface{}) (err error) {
	if s.top+1 > len(s.data) {
		return errors.New("stack overload error")
	}
	s.top++
	s.data[s.top] = e
	return nil
}

func (s *MyStack) Pull() (e interface{}) {
	if s.top < 0 {
		return nil
	}
	e = s.data[s.top]
	s.data[s.top] = nil
	s.top--
	return
}

func (s *MyStack) GetTop() (e interface{}) {
	if s.top == -1 {
		return nil
	}
	return s.data[s.top]
}
