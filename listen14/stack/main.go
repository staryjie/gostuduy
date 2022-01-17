package main

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

func main() {
	var s Stack
	fmt.Println(s.Empty())
	s.Push("hello")
	fmt.Println(s.Empty())
	//ret, _ := s.Pop()
	//fmt.Println(ret)
	fmt.Println(s.Empty())
	s.Push("world")
	ret, _ := s.Top()
	fmt.Println(ret)

	for _, e := range s.data {
		if e != "" {
			fmt.Println(e)
		}
	}
}
