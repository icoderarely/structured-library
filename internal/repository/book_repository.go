package repository

import (
	"errors"
	"sync"

	"github.com/icoderarely/structured-library/internal/domain"
)

type inMemBookRepository struct {
	mu sync.RWMutex
	db map[int]*domain.Book
	ID int
}

func NewBookRepository() domain.BookRepository {
	return &inMemBookRepository{
		db: make(map[int]*domain.Book),
	}
}

func (r *inMemBookRepository) Create(book *domain.Book) (*domain.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ID++
	book.ID = r.ID
	r.db[r.ID] = book
	return book, nil
}

func (r *inMemBookRepository) GetById(ID int) (*domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	book, ok := r.db[ID]
	if !ok {
		return nil, errors.New("book not found")
	}

	return book, nil
}

func (r *inMemBookRepository) List() ([]*domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	books := make([]*domain.Book, 0)
	for _, book := range r.db {
		books = append(books, book)
	}

	return books, nil
}

func (r *inMemBookRepository) Delete(ID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.db[ID]
	if !ok {
		return errors.New("book not found")
	}

	delete(r.db, ID)

	return nil
}
