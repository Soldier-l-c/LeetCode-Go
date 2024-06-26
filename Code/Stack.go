package Code

type Stack struct {
	size  int
	array []interface{}
}

func NewStack() *Stack {
	s := new(Stack)
	s.array = []interface{}{}
	s.size = 0
	return s
}

func (s *Stack) Push(v interface{}) {
	s.array = append(s.array, v)
	s.size++
}

func (s *Stack) Top() interface{} {
	if s.Empty() {
		return nil
	}
	return s.array[s.size-1]
}

func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	res := s.array[s.size-1]
	s.array = s.array[0 : s.size-1]
	s.size--
	return res
}

func (s *Stack) Empty() bool {
	return s.size == 0
}
