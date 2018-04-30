package iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {
	bookShelf := NewBookShelfSetMaxSize(3)
	asserts := []string{"A", "B", "C"}
	for _, asserts := range asserts {
		if err := bookShelf.Add(&Book{asserts}); err != nil {
			t.Errorf("Expect not error: %s", err)
		}
	}
	i := 0
	for iterator := bookShelf.Iterator(); iterator.HasNext(); {
		if book := iterator.Next(); book.(*Book).GetName() != asserts[i] {
			t.Errorf("Expect Book.name to %s, but %s", asserts[i], book.(*Book).GetName())
		}
		i++
	}
}
