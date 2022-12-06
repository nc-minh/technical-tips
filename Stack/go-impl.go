package main

type Stack struct {
	elements []string
}

func (s *Stack) Push(e string) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	e := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return e, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Top() string {
	return s.elements[len(s.elements)-1]
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func main() {
	var s Stack
	s.Push("hi")
	s.Push("wow")
	s.Push("land")

	println(s.Top())
	println(s.Size())
}
