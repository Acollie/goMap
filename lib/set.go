package lib

type Set[K comparable] struct {
	items map[K]bool
	len   int
}

func (set *Set[K]) New() *Set[K] {
	set.items = make(map[K]bool)
	return set
}
func (set *Set[K]) Add(item K) {
	set.items[item] = true
	set.len += 1
}
func (set *Set[K]) Remove(item K) {
	_, ok := set.items[item]
	if ok {
		delete(set.items, item)
		set.len -= 1
	}
}
func (set *Set[K]) Len() int {
	return set.len
}
func (set *Set[K]) In(item K) bool {
	_, ok := set.items[item]
	if ok {
		return true
	} else {
		return false
	}
}
