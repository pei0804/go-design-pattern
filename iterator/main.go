package iterator

import "errors"

// Aggregate 集合体を表すインターフェイス
type Aggregate interface {
	Iterator() Iterator
}

// Iterator 数え上げ、スキャンを行うインターフェイス
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// BookShelf 本棚
type BookShelf struct {
	Books   []*Book
	maxSize int
}

// NewBookShelfSetMaxSize 本棚サイズの上限を指定して作成
func NewBookShelfSetMaxSize(maxSize int) *BookShelf {
	return &BookShelf{
		Books:   []*Book{},
		maxSize: maxSize,
	}
}

// NewBookShelf 本棚サイズの上限なしで作成
func NewBookShelf() *BookShelf {
	return &BookShelf{
		Books: []*Book{},
	}
}

// Add 本を追加する
func (bs *BookShelf) Add(book *Book) error {
	if bs.maxSize != 0 && len(bs.Books) >= bs.maxSize {
		return errors.New("本棚の上限を越えています")
	}
	bs.Books = append(bs.Books, book)
	return nil
}

// Iterator 集合を示す
func (bs *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{BookShelf: bs}
}

// BookShelfIterator 本棚スキャン
type BookShelfIterator struct {
	BookShelf *BookShelf
	last      int
}

// HasNext 次に要素が存在するか
func (bsi *BookShelfIterator) HasNext() bool {
	if bsi.last >= len(bsi.BookShelf.Books) {
		return false
	}
	return true
}

// Next 次の要素を取得する
func (bsi *BookShelfIterator) Next() interface{} {
	book := bsi.BookShelf.Books[bsi.last]
	bsi.last++
	return book
}

// Book 本
type Book struct {
	name string
}

// GetName 本の名前を取得する
func (b *Book) GetName() string {
	return b.name
}
