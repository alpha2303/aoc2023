package utils

/*
	Set - Generic HashSet Implementation in Go using in-built Map datatype
*/
type Set[K comparable] struct {
	list map[K]struct{}
}

func (s *Set[K]) Add(val K) {
	s.list[val] = struct{}{}
}

func (s *Set[K]) Has(val K) bool {
	_, ok := s.list[val]
	return ok
}

func (s *Set[K]) Remove(val K) {
	delete(s.list, val)
}
