package generics


type StackOfInts struct {
	elements []int
}

func (s *StackOfInts) Push(element int) {
	s.elements = append(s.elements, element)
}

func (s *StackOfInts) IsEmpty() bool {
	return len(s.elements) == 0
}		

func (s *StackOfInts) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return element, true
}


type StackOfStrings struct {
	elements []string
}

func (s *StackOfStrings) Push(value string) {
	s.elements = append(s.elements, value)
}

func (s *StackOfStrings) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *StackOfStrings) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(s.elements) - 1
	el := s.elements[index]
	s.elements = s.elements[:index]
	return el, true
}


type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}		

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return element, true
}
