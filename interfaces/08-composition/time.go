package main

// usually: include the time method in the list.go
// it is here for clarity

// TODO: NEW
// you could include a time method in the book and game instead.
// but sometimes it is not possible to do so.
func (l list) time() (total int) {
	for _, it := range l {
		switch it := it.(type) {
		case *game:
			total += it.playTime
		case book:
			total += it.readTime
		}
	}
	return total
}
