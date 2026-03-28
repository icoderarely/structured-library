package service

import (
	"errors"

	"github.com/icoderarely/structured-library/internal/domain"
)

type BookService struct {
	repo domain.BookRepository
}

func NewBookService(repo domain.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(book *domain.Book) (*domain.Book, error) {
	if book.Author == "" || book.Name == "" {
		return nil, errors.New("author and/or title can't be emplty")
	}
	return s.repo.Create(book)
}

func (s *BookService) GetBook(id int) (*domain.Book, error) {
	return s.repo.GetById(id)
}

func (s *BookService) GetBooks() ([]*domain.Book, error) {
	return s.repo.List()
}

func (s *BookService) DeleteBook(id int) error {
	return s.repo.Delete(id)
}
