package main

import "math"

type MinStack struct {
	arr []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int)  {
	s.arr = append(s.arr, x)
}

func (s *MinStack) Pop()  {
	s.arr = append(make([]int,0), s.arr[:len(s.arr)-1]...)
}

func (s *MinStack) Top() int {
	return s.arr[len(s.arr)-1]
}

func (s *MinStack) GetMin() int {
	min := math.MaxInt64
	for _, n := range s.arr {
		if n < min {
			min = n
		}
	}
	return min
}
