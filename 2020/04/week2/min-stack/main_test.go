package main

import "testing"

func TestMinStack(t *testing.T) {
	s := Constructor()

	s.Push(1)
	s.Push(2)
	s.Push(-1)
	s.Push(0)

	var v int

	for i := 1; i <= 2; i++ {
		v = s.GetMin()
		if v != -1 {
			t.Fatalf("#%d - expected -1 but was %d", i, v)
		}

		v = s.Top()
		if v != 0 {
			t.Fatalf("#%d - expected 0 but was %d", i, v)
		}
	}

	s.Pop()
	v = s.GetMin()
	if v != -1 {
		t.Fatalf("expected -1 but was %d", v)
	}
	v = s.Top()
	if v != -1 {
		t.Fatalf("expected -1 but was %d", v)
	}

	s.Pop()
	v = s.GetMin()
	if v != 1 {
		t.Fatalf("expected 1 but was %d", v)
	}
	v = s.Top()
	if v != 2 {
		t.Fatalf("expected 2 but was %d", v)
	}
}
