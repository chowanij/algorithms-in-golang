package collections

// Item the type of the Set
type Item interface{}

// JCSet the set of Items
type JCSet struct {
	uniques map[Item]bool
	items   []*Item
}

// Add adds a new element to the Set. Returns a pointer to the Set.
func (s *JCSet) Add(t Item) *JCSet {
	if s.uniques == nil {
		s.uniques = make(map[Item]bool)
	}
	_, ok := s.uniques[t]
	if !ok {
		s.uniques[t] = true
		s.items = append(s.items, &t)
	}
	return s
}

// Pop - mimics Pop method from Python Set implementation ;)
func (s *JCSet) Pop() Item {
	var i Item
	// l := len(s.items) - 1
	// indx := 0
	if s.emptyOrUnInit() {
		return nil
	}

	i = *s.items[0]
	s.removeIndex(0)
	delete(s.uniques, i)

	return i
}

// Remove - delete element from set
func (s *JCSet) Remove(i Item) Item {
	if s.emptyOrUnInit() {
		return nil
	}

	item := s.uniques[i]
	delete(s.uniques, i)

	return item
}

// check if set is initialized or is not empty
func (s *JCSet) emptyOrUnInit() bool {
	return s.uniques == nil || len(s.uniques) == 0
}

func (s *JCSet) removeIndex(index int) {
	s.items = append(s.items[:index], s.items[index+1:]...)
}
