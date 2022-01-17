package utils

import "fmt"

type Stack struct {
	data [1024]string
	top  int
}

func (s *Stack) Push(d string) {
	s.data[s.top] = d
	s.top++
}

func (s *Stack) Pop() (ret string, err error) {
	if s.top == 0 {
		err = fmt.Errorf("Stack is empty!")
		return
	}
	s.top--
	ret = s.data[s.top]
	return
}

func (s *Stack) Top() (ret string, err error) {
	if s.top == 0 {
		err = fmt.Errorf("Stack is empty!")
		return
	}
	ret = s.data[s.top-1]
	return
}

func (s *Stack) Empty() bool {
	return s.top == 0
}
