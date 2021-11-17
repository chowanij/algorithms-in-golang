package collections

import (
	"fmt"
)

// Item the type of the Set
type Item interface{}

// JCSet the set of Items
type JCSet struct {
	uniques map[Item]bool
	items   []*Item
}

// Items - retrieve all items from set
func (s *JCSet) Items() []Item {
	deref := make([]Item, len(s.items), cap(s.items))
	for idx, item := range s.items {
		deref[idx] = *item
	}

	return deref
}

// Equals - compares 2 sets and return tru if equls, flase otherwise
func (s JCSet) Equals(s2 JCSet) bool {
	if len(s.items) != len(s2.items) {
		return false
	}
	
	for _, i := range s.Items() {
		_, ok := s2.uniques[i]
		if !ok {
			return false
		}
	}

	return true
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
	if s.emptyOrNotInit() {
		return nil
	}

	i = *s.items[0]
	s.removeIndex(0)
	delete(s.uniques, i)

	return i
}

// Remove - delete element from set
func (s *JCSet) Remove(i Item) Item {
	if s.emptyOrNotInit() {
		return nil
	}

	item := s.uniques[i]
	delete(s.uniques, i)

	return item
}

// Clear - remove all items from set
func (s *JCSet) Clear() {
	for k := range s.uniques {
		delete(s.uniques, k)
	}
	s.items = make([]*Item, 0)
}

// Intersection - returns sets intesetction as set
func (s *JCSet) Intersection(s2 *JCSet) JCSet {
	var s3 JCSet

	for i := range (*s).uniques {
		_, ok := (*s2).uniques[i]
		if ok {
			s3.Add(i)
		}

	}

	return s3
}

func (s JCSet) String() string {
	deref := make([]Item, len(s.items), cap(s.items))

	for idx, item := range s.items {
		deref[idx] = *item
	}

	return fmt.Sprintf("%v", deref)
}

// check if set is initialized or is not empty
func (s *JCSet) emptyOrNotInit() bool {
	return s.uniques == nil || len(s.uniques) == 0
}

func (s *JCSet) removeIndex(index int) {
	s.items = append(s.items[:index], s.items[index+1:]...)
}
