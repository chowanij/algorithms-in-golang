package collections

// Item the type of the Set
type Item interface{}

// JCSet the set of Items
type JCSet struct {
	items map[Item]bool
}

// Add adds a new element to the Set. Returns a pointer to the Set.
func (s *JCSet) Add(t Item) *JCSet {
	if s.items == nil {
		s.items = make(map[Item]bool)
	}
	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
	return s
}

// Pop - mimics Pop method from Python Set implementation ;)
func (s *JCSet) Pop() Item {
	var i Item
	l := len(s.items) - 1
	indx := 0
	if s.emptyOrUnInit() {
		return nil
	}

	// maps in GO behave like LIFO
	// so in order to make Pop() behaves like in Python we need take out last element
	// revers order of poping elements
	for k := range s.items {
		if indx == l {
			i = k
		}
		indx++
	}

	delete(s.items, i)

	return i
}

// Remove - delete element from set 
func (s *JCSet) Remove(i Item) Item {
	if s.emptyOrUnInit() {
		return nil
	}

	item := s.items[i]
	delete(s.items, i)

	return item
}

// check if set is initialized or is not empty
func (s *JCSet) emptyOrUnInit() bool {
	return s.items == nil || len(s.items) == 0
}
